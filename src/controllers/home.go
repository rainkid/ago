package controllers

import (
	"io"
	"net/http"
	"html/template"
	"crypto/md5"
    "strconv"
    "time"
    "fmt"
)
func Home(w http.ResponseWriter, r *http.Request) {
	crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))

    t, _ := template.ParseFiles("src/views/index/index.html")
    t.Execute(w, token)
}
