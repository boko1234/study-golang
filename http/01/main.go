package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	hotpot()
}

func hotpot() {
	url := "http://localhost:4000/post"
	contentType := "application/json"
	data := `{"name":"小王子","age":17}"`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
	    fmt.Println("errrrr",err)
		return
	}
	defer resp.Body.Close()
	b,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
	    fmt.Println("get resp failed,err:%v\n",err)
		return
	}
	fmt.Println(string(b))
}

func hotgot() {
	resp, err := http.Get("http://www.xvwvwv.top")
	if err != nil {
		fmt.Println("errrrrrrrrrr", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read form resp.body failed,err", err)
		return
	}
	fmt.Println(string(body))
}
