package utils

import (
	"crypto/md5"
	"dogo"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

func GetConfig(filename string, name string) string {
	config, err := dogo.NewConfig(fmt.Sprintf("src/configs/%s.yaml", filename))
	if err != nil {
		return ""
	}
	str, err := config.String(ENV(), name)
	if err != nil {
		return ""
	}
	return str
}

func ENV() string {
	config, err := dogo.NewConfig("src/configs/app.yaml")
	if err != nil {
		return "product"
	}
	env, err := config.String("base", "env")
	if err != nil {
		return "product"
	}
	return env
}

func MD5(str string) []byte {
	h := md5.New()
	io.WriteString(h, str)
	return h.Sum(nil)
}

func IValue(i interface{}) (value []byte, length int64) {
	if i != nil {
		rt := reflect.TypeOf(i)
		rv := reflect.ValueOf(i)
		switch rt.Kind() {
		case reflect.Int:
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
			value = []byte(strconv.FormatInt(rv.Int(), 10))
		case reflect.Uint:
		case reflect.Uint8:
		case reflect.Uint16:
		case reflect.Uint32:
		case reflect.Uint64:
			value = []byte(strconv.FormatUint(rv.Uint(), 10))
		case reflect.Float32:
		case reflect.Float64:
			value = []byte(strconv.FormatFloat(rv.Float(), 'f', -1, 64))
		case reflect.Slice:
			if rt.Elem().Kind() == reflect.Uint8 {
				value = rv.Interface().([]byte)
				break
			}
		case reflect.String:
			value = []byte(rv.String())
		default:
			value = []byte(rv.String())
		}
	}
	return value, int64(len(value))
}
