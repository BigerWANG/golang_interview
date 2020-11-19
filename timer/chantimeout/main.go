package main

import (
	"time"
	"fmt"
)

/*
超时定时器


在time.Time中，对外通知定时器到期的途径就是通道，由字段C代表
C是一个chan time.Time 类型的带缓冲的接收通道

*/

func timebase() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Present time： %v. \n", time.Now())

	expirationTime := <- timer.C
	fmt.Printf("Expiration time: %v \n", expirationTime)

	// Reset方法用于重置定时器（也就是说，定时器是可以复用的）也会返回一个bool类型值
	timer.Reset(time.Second*3)
	expirationTime1 := <- timer.C
	fmt.Printf("Expiration time: %v \n", expirationTime1)

	// Stop 方法用于停止计时器，返回一个bool类型值
	fmt.Printf("Stop timer: %v. ", timer.Stop())
}



func sampleTimeoutTimer(){
	// 简单的定时器
	intChan := make(chan int, 1)
	go func() {  // 一个worker
		time.Sleep(time.Millisecond * 200)
		intChan <- 1
	}()

	select {
	case i := <- intChan:
		fmt.Printf("Received: %d.\n", i)
		//case <-time.After(time.Millisecond * 500):

	case <-time.NewTimer(time.Millisecond * 500).C:  // 每次有需要新生成一个timer对象
		fmt.Println("Timeout! ")
	}
}

func advtangeTimeoutTimer()  {
	// 升级版定时器，实现了对计时器的复用
	var timer *time.Timer

	intChan := make(chan int, 1)
	timeout := time.Millisecond * 500

	go func() { // 模拟一个worker
		for i:=0; i<5; i++{
			time.Sleep(time.Second*1)
			intChan<-i
		}
		close(intChan)
	}()


	for{
		if timer == nil{
			timer = time.NewTimer(timeout)
		}else{
			timer.Reset(timeout) // 每次复用同一个计时器
		}

		select{
		case e, ok := <-intChan:
			if !ok{
				fmt.Println("End. ")
				timer.Stop()
				return
			}
			fmt.Printf("Received: %v\n", e)
		case <-timer.C:  // 等待计时器...
			fmt.Println("TimeOut!")
		}
	}
	
}


func main()  {
	timebase()

	sampleTimeoutTimer()

	advtangeTimeoutTimer()

}
