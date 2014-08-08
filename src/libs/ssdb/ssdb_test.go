package ssdb

import (
	"testing"
)

func Test_Conn(t *testing.T) {
	db, err := Connect("127.0.0.1", 8888)
	if err != nil {
		t.Error(err)
	}
	db.Close()

	/*keys := []string{}
	keys = append(keys, "c")
	keys = append(keys, "d")
	val, err = db.Do("multi_get", "a", "b", keys)
	fmt.Printf("%s\n", val)


	db.Del("a")
	val, err = db.Get("a")
	fmt.Printf("%s\n", val)

	db.Do("zset", "z", "a", 3)
	db.Do("multi_zset", "z", "b", -2, "c", 5, "d", 3)
	resp, err := db.Do("zrange", "z", 0, 10)
	if err != nil {
		t.Error(err)
	}
	if len(resp)%2 != 1 {
		fmt.Printf("bad response")
		t.Error(err)
	}

	fmt.Printf("Status: %s\n", resp[0])
	for i := 1; i < len(resp); i += 2 {
		fmt.Printf("  %s : %3s\n", resp[i], resp[i+1])
	}
	return*/
}

func Test_Set(t *testing.T) {
	db, err := Connect("127.0.0.1", 8888)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = db.Set("a", "xxx")
	if err != nil {
		t.Error(err.Error())
	}
}

func Test_Get(t *testing.T) {
	db, err := Connect("127.0.0.1", 8888)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	val, err := db.Get("a")
	if err != nil {
		t.Error(err.Error())
	}
	if val != "xxx" {
		t.Error("get value not setted.")
	}
}

func Test_Del(t *testing.T) {
	db, err := Connect("127.0.0.1", 8888)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = db.Del("a")
	if err != nil {
		t.Error(err.Error())
	}
}
