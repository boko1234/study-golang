package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)


var dor, _ = os.Getwd()

func main() {
	//dir, _ := os.Executable()
	//fmt.Println(os.Getwd())
	//fmt.Println(filepath.Dir(dor))
	a := ":5656"
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/hello", dddas)
	fmt.Println("http://127.0.0.1" + a)
	fmt.Println("http://localhost" + a)
	err := http.ListenAndServe(a, nil)
	if err != nil {
		fmt.Println("errrrrrrrrrr", nil)
		return
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(dor+"/golang/http/local/home.tmpl")
	if err != nil {
		fmt.Println("create template failed, err :",err)
		return
	}
	_ = tmpl.Execute(w, "home")
}
func dddas(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(dor+"/golang/http/local/hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err :",err)
		return
	}
	_ = tmpl.Execute(w, "hello")
}
