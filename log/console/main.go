package main

import (
	"../mylogger"
	"time"
)

func main() {
	for {
		log := mylogger.NewLog()
		log.Debug("这是一条Debug")
		log.Info("这是一条Info")

		time.Sleep(time.Second)
	}

}
