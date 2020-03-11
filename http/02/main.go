package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	a := ":5678"
	fmt.Println(a)
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/post", dddas)
	fmt.Println("http://127.0.0.1" + a)
	err := http.ListenAndServe(a, nil)
	if err != nil {
		fmt.Println("errrrrrrrrrr", nil)
		return
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `欢迎来到主页<a href="/hello">hello页面</a>`)

}
func dddas(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println("1",r.PostForm) // 打印form数据
	fmt.Println("2",r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println("3",string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))

}
