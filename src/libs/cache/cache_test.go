package cache

import (
	"fmt"
	"testing"
	"time"
)

var (
	acache = NewCaches()
)

func Test_Prof(t *testing.T) {
	ori := time.Now()
	var i int

	for i = 0; i <= 10000; i++ {
		acache.Set(fmt.Sprintf("key_%d", i), i, (int64)(10+i))
	}

	fmt.Println("list: " + time.Now().Sub(ori).String())
}
