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
	config, err := dogo.NewConfig(fmt.Sprintf("src/configs/%s.ini", filename))
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
	config, err := dogo.NewConfig("src/configs/app.ini")
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

/*func IValue(i interface{}) (value []byte, length int64) {
	if i != nil {
		rt := reflect.TypeOf(i)
		rv := reflect.ValueOf(i)
		switch rt.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = []byte(strconv.FormatInt(rv.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			value = []byte(strconv.FormatUint(rv.Uint(), 10))
		case reflect.Float32:
		case reflect.Float64:
			value = []byte(strconv.FormatFloat(rv.Float(), 'f', -1, 64))
		case reflect.Slice, reflect.Array:
			switch rt.Elem().Kind() {
			case reflect.Uint8:
				value = rv.Interface().([]byte)
			case reflect.String:
				value = []byte(reflect.ValueOf(rv.Interface()).Index(0).String())
			default:
				value = nil
			}
		case reflect.String:
			value = []byte(rv.String())
		default:
			value = nil
		}
	}
	return value, int64(len(value))
}*/
