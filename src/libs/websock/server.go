package websock

import(
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
)

type SockServer struct{
	client *websocket.Conn
	data chan([]byte)
}
var  (
	Server = NewServer()
)

func NewServer() *SockServer{
	return &SockServer{
		data :make(chan []byte),
	}
}

func Start() {
	http.Handle("/sock", websocket.Handler(Server.Run))
	fmt.Println("/sock start for websocket");
}

func (m *SockServer) Run(ws *websocket.Conn){
	defer func() {
		if err := ws.Close(); err != nil {
			fmt.Println("Websocket could not be closed", err.Error())
		}
	}()
	m.client = ws
	fmt.Println(ws.Request().RemoteAddr, " connected.",)
	for {
		select {
			case data := <-m.data:
			m.send(data)
			break
		}
	}
}
func (m *SockServer) SendData(cmd []byte) {
	m.data<-cmd
}

func (m *SockServer) send(cmd []byte){
	if m.client != nil {
		if err := websocket.JSON.Send(m.client, cmd); err != nil {
			// we could not send the message to a peer
			fmt.Println("Could not send message ", err.Error())
		}
	}
}
