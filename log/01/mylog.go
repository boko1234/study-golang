package _1

import "fmt"

type Logger struct {

}

func NewLog() Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	fmt.Println(msg)
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}
func (l Logger) Warning(msg string) {
	fmt.Println(msg)
}
func (l Logger) Error(msg string) {
	fmt.Println(msg)
}

func (l Logger) Fatal(msg string) {
	fmt.Println(msg)
}
