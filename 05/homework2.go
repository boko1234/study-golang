package main

import "fmt"

//学生管理系统
//有一个物件
//1.它保存了一些数据 ===》结构体的字段
//2.它有4个功能 ---》结构体的方法

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudent map[int64]student
}

func (s studentMgr) showStudents() {
	for _, stu := range s.allStudent {
		fmt.Println(stu.id, stu.name)
	}
}
func (s studentMgr) addStudents() {
	var (
		stuId   int64
		stuName string
	)
	fmt.Print("请输入学号")
	fmt.Scanln(&stuId)
	fmt.Print("请输入姓名")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuId,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
}
func (s studentMgr) exitStudents() {

}
func (s studentMgr) delStudents() {

}

func main() {

}
