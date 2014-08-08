package utils

import (
	"fmt"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	var str = "1234567890"
	estr, err := Encrypt(str, "12345678")
	if err != nil {
		t.Error("encrypt error.")
	}

	dstr, err := Decrypt(estr, "12345678")
	if err != nil {
		t.Error("encrypt error.")
	}

	if dstr != str {
		t.Error("decrypt error")
	}
	fmt.Println(dstr)
}
