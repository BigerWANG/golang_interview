package main

import (
	"fmt"
	//"sync"
	"time"
)




func PrioritySelect(ch1, ch2 <- chan int, stopCh chan struct{})  {
	dojob := func(jobnum int){
		fmt.Println("do job...", jobnum)
		time.Sleep(time.Microsecond * 100)
		fmt.Printf("job %d is done\n", jobnum)
	}

	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			dojob(job1)
		case job2 := <- ch2:
			priority:
				for {
					select {
					case job1 := <- ch1:
						dojob(job1)
					default:
						break priority
					}
				}
				dojob(job2)
		}
	}

}

// 带注释版
func worker(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Printf("job%d done\n", job1)
		case job2 := <-ch2:  // 读取到ch2中有数据
			fmt.Println("job2 is ready...")
		priority:  // 设置 LABEL
			for {  // 再去检查ch1
				fmt.Println("check job1 again...")
				select {
				case job1 := <-ch1:
					fmt.Printf("job%d done\n", job1)
				default:
					fmt.Println("job1 also nil, break to exec job2...")
					break priority  // 如果不存在退出此次select 到 priority锚点
				}
			}
			fmt.Printf("job%d done\n", job2)  // 最终执行job2...
		}
		fmt.Println("Im waiting for...")
	}
}





func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	stop := make(chan struct{}, 1)
	go func() {
		ch1 <- 1
		stop <- struct{}{}
	}()

	go func() {
		ch2 <- 2
	}()

	PrioritySelect(ch1, ch2, stop)
}