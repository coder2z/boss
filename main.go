package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"time"
)

func main() {

	fmt.Printf("开始：按q结束 : %s \n", time.Now())
	fmt.Printf("5秒准备时间==> \n")
	tick := time.Tick(1 * time.Second)
	for countdown := 5; countdown > 0; countdown-- {
		fmt.Printf("\r%2d", countdown)
		<-tick
	}
	go func() {
		for {
			if robotgo.AddEvent("q") {
				fmt.Println("exit")
				os.Exit(-1)
			}
		}
	}()

	for {
		fmt.Printf("运行中：%s \n", time.Now())
		robotgo.KeyTap("right")
		robotgo.Click()
		time.Sleep(time.Second)
	}
}
