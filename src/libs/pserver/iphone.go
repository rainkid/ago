package pserver

import (
	"bytes"
	"fmt"
	"net"
	"libs/websock"
)

var (
	Server *PServer
)

type Client struct {
	data []byte
	conn net.Conn
}

type PServer struct {
	listener net.Listener
	clients  map[string]*Client
	qmsg     chan *Client
	qclose   chan *Client
	err      error
}

func (client *Client) Close() {
	client.conn.Close()
}

func (client *Client) Recv() {
	for {
		buf := make([]byte, 512)
		bl, err := client.conn.Read(buf)

		if err != nil && err.Error() == "EOF" {
			Server.qclose <- client
			break
		}

		if len(buf) > 0 {
			msgs := bytes.Split(buf[0:bl], []byte("\r\n"))
			for _, msg := range msgs {
				if len(msg) > 0 {
					client.data = msg
					Server.qmsg <- client
				}
			}

		}
	}
}

func (client *Client) Send(buf []byte) {
	cmd := string(buf) + "\r\n"
	fmt.Println("send data to " + client.conn.RemoteAddr().String() + " " + cmd)
	client.conn.Write([]byte(cmd))
}

func NewPServer() *PServer {
	return &PServer{
		qmsg:    make(chan *Client),
		qclose:  make(chan *Client),
		clients: make(map[string]*Client),
	}
}

func Start() *PServer {
	if Server == nil {
		Server = NewPServer()
		go Server.Listen()
		go Server.Daemon()
	}
	return Server
}

func (i *PServer) Daemon() {
	go func() {
		for {
			select {
			case client := <-i.qmsg:
				i.DealMsg(client)
				break
			case client := <-i.qclose:
				i.DelClient(client)
				break
			}
		}
	}()
}

func (i *PServer) SendData(buf []byte) {
	for _, client := range i.clients {
		client.Send(buf)
	}
}

func (i *PServer) DealMsg(client *Client) {
	cmd := string(client.data)
	switch cmd {
	case "exit":
		i.DelClient(client)
		break
	default:
		fmt.Println(cmd)
		websock.Server.SendData(client.data);
		break
	}
}

func (i *PServer) AddClient(client *Client) {
	go client.Recv()
	mkey := client.conn.RemoteAddr().String()
	fmt.Println(mkey + " connected.")
	i.clients[mkey] = client
}

func (i *PServer) DelClient(client *Client) {
	mkey := client.conn.RemoteAddr().String()
	i.clients[mkey].Close()
	delete(i.clients, mkey)
	fmt.Println(client.conn.RemoteAddr().String() + " closed")
}

//开始服务器
func (i *PServer) Listen() {
	listener, err := net.Listen("tcp", "0.0.0.0:8008")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("listen " + listener.Addr().String())
	i.listener = listener
	for {
		conn, err := i.listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		client := &Client{conn: conn}
		i.AddClient(client)
	}
}

func (i *PServer) Close() {
	for _, client := range i.clients {
		client.Close()
	}
	i.listener.Close()
}
