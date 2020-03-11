package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// 在终端写日志相关内容

type LogLevel uint16

const (
	UNKNON LogLevel = iota //DEBUG定义日志级别
	TRACE
	DEBUG
	INFO
	WAENING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WAENING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNON, err
	}
}

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level:level,
	}
}

func (l Logger) Debug(msg string) {
	if l.Level > DEBUG {
		now := time.Now()
		fmt.Printf("[%s] [DEBUG]  %s", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Info(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [INFO] %s", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Warning(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [WARNING] %s", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Error(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [ERROR] %s", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Fatal(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [FATAL] %s", now.Format("2006-01-02 15:04:05"), msg)
}
