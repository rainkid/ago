package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/rainkid/dogo"
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

var (
	loger = log.New(os.Stdout, "[libs] ", log.Ldate|log.Ltime)
	env   = "product"
)

func GetConfig(filename string, name string) string {
	cfgpath := dogo.Register.Get("cfg_path")
	config, err := dogo.NewConfig(fmt.Sprintf("%s/%s.ini", cfgpath, filename))
	if err != nil {
		return ""
	}
	str, err := config.String(Env(), name)
	if err != nil {
		return ""
	}
	return str
}

func Env() string {
	inipath := dogo.Register.Get("app_ini")
	config, err := dogo.NewConfig(inipath)
	if err != nil {
		loger.Print(err.Error())
		return env
	}
	env, err := config.String("base", "environ")
	if err != nil {
		loger.Print(err.Error())
		return env
	}
	return env
}

func MD5(str string) []byte {
	h := md5.New()
	io.WriteString(h, str)
	return h.Sum(nil)
}

func ItoByte(i interface{}) []byte {
	return []byte(ItoString(i))
}

func ItoString(i interface{}) string {
	rt := reflect.TypeOf(i)
	switch rt.Kind() {
	case reflect.String:
		return reflect.ValueOf(i).Index(0).String()
	case reflect.Slice:
		switch rt.Elem().Kind() {
		case reflect.String:
			return reflect.ValueOf(i).Index(0).String()
		case reflect.Uint8:
			return fmt.Sprintf("%s", i)
		}
		return fmt.Sprintf("%s", i)
	}
	return fmt.Sprintf("%s", i)
}

func ItoInt(i interface{}) int64 {
	ret, _ := strconv.Atoi(ItoString(i))
	return int64(ret)
}

func ItoStrings(i interface{}) []string {
	len := reflect.ValueOf(i).Len()
	var strs []string
	var j int
	for j = 0; j < len; j++ {
		strs = append(strs, reflect.ValueOf(i).Index(j).String())
	}
	return strs
}

func RandString(len int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < len; {
		if string(RandInt(65, 90)) != temp {
			temp = string(RandInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func MapMerge(a, b []interface{}) []interface{} {
	var c []interface{}
	for _, value := range b {
		c = append(a, value)
	}
	return c
}
