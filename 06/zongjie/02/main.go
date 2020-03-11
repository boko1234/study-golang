package main

import (
	"fmt"
	"io"
	"os"
)

func f1() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	if err != nil {
		fmt.Println(err, "errrrrrrrrrrrr")
	}
	defer fileObj.Close()
}

func f2() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("errrr", err)
		return
	}

	tmpFile,err:= os.OpenFile("./sb.tmp",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0644)
	if err != nil {
		fmt.Println("ererer",err)
		return
	}
	defer tmpFile.Close()
	//读原文件写入临时文件

	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Println("errrrrrreeee", err)
		return
	}


	tmpFile.Write(ret[:n])
	//在写入原文件

	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)

	var x [1024]byte
	for{
		n, err := fileObj.Read(x[:])
		if err != io.EOF {
			break
		}
		if err != nil {
			fmt.Println("EEEEErrr",err)
			return
		}
		tmpFile.Write(x[:n])
	}
	//原文件后续也写入了临时文件中
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp","./sb.txt")





}

func main() {
	f2()
}
