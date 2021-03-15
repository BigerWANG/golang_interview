package main
import (
	"fmt"
	"time"
)

func main() {
	// 定时器，定时向通道中发送消息
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)  // new一个定时器，设定一个轮训间隔时间
	go func() {
		for range ticker.C{  // 表示设置的时间到期了
			select {
			case intChan <- 1:  // 表示随机发送
			case intChan <- 2:
			case intChan <- 3:
			}
		}
		fmt.Printf("End. [sender]")
	}()

	var sum int

	for e:=range intChan{
		fmt.Printf("Received: %v\n", e)
		sum+=e
		if sum > 10{
			fmt.Printf("Got: %v\n", sum)
			ticker.Stop()
			break
		}
		fmt.Println("End. [receiver]")
	}
}