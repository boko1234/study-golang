package main

import (
	"fmt"
	"os"
)

var (
	allStudent map[int64]*student
)

type student struct {
	id   int64
	name string
}
//新建学生信息
func newStudent(id int64, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)

	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	newStu := newStudent(id,name)
	allStudent[id] = newStu
	//allStudent[id] = newStu
}
func delStudent() {
	var (
		deleteId int64
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&deleteId)
	delete(allStudent,deleteId)
}

func main() {
	allStudent = make(map[int64]*student, 40)
	for {
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(`
		1查看所有学生
		2.新增学生
		3.删除学生
	`)

		fmt.Print("请输入你要干啥")

		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项：、n", choice)

		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("GUN~")
		}
	}

}
