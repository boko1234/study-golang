package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func lotgot(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request .body failed,err",err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}
func main() {
	a := ":5678"
	fmt.Println(a)
	http.HandleFunc("/", lotgot)
	//http.HandleFunc("/hello", dddas)
	fmt.Println("http://127.0.0.1"+a)
	err := http.ListenAndServe(a, nil)
	if err != nil {
		fmt.Println("errrrrrrrrrr", nil)
		return
	}
}
