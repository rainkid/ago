package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/rainkid/dogo"
	"io"
	"reflect"
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

func ItoByte(i interface{}) []byte {
	return []byte(fmt.Sprintf("%s", i))
}

func ItoString(i interface{}) string {
	return reflect.ValueOf(i).Index(0).String()
}

func ItoUint(i interface{}) uint64 {
	return reflect.ValueOf(i).Uint()
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
