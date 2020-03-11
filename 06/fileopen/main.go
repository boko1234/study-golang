package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readfile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed,", err)
		return
	}
	defer fileObj.Close()

	//var tmp = make([]byte,128)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
		fmt.Println(string(tmp[:n]))
		if n == 0 {
			return
		}
	}
}

func readfileFrom() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("errererererer", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("errrrrrrrrrr", err)
			return
		}
		fmt.Print(line)
	}
}

func readFF() {
	fileObj, err := os.OpenFile("./../../data.txt", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
	//file, err := os.OpenFile("./../../data.json",, 0644)
	if err != nil {
		fmt.Println("errrr", err)
		return
	}
	_, _ = fileObj.Write([]byte("eeeeeeeeeeeeeeeeeeee"))
	_, _ = fileObj.WriteString("aaaaaaaaaaaaaaaaaa")
	fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./../../aa.txt", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("errrr", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	_, _ = wr.WriteString("helloSHAHE")
	_ = wr.Flush()
}





func main() {
	writeDemo2()
}
