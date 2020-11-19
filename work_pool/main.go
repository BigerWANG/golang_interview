package main

import (
	"math/rand"
	"time"
	"fmt"
)

/*
使用goroutine 和channel实现一个计算int64随机数各位数和的程序
123 -> 6
345 -> 12

1，开启一个goroutine 循环生成int64的随机数，发送到jobChan
2, 开启24个goroutine从jobChan中取出随机数计算各位的和，将结果发送到resultChan
3, 主goroutine 从resultChan中取结果并打印到终端输出
*/

type job struct {
	jobNum int64
}

type result struct {
	job *job
	sum int64
}


func cal(num int64) int64{
	/*
	计算int64每个位数相加的总和
	*/
	var sum int64
	for num >= 1{
		ret := num % 10
		sum += ret
		num /= 10
	}
	return sum
}


func genJob(jchan chan<- *job)  {
	for {
		jchan <- &job{
			jobNum: rand.Int63(),
		}
		time.Sleep(time.Second* 1)
	}
	
}

func getJob2Res(jchan <-chan *job, reschan chan <- *result){
	for job := range jchan{
		reschan <- &result{
			job: job,
			sum: cal(job.jobNum),
		}
	}

}


func main() {
	jobChan := make(chan *job)
	resultChan := make(chan *result)

	go genJob(jobChan)

	for i:=0; i<=24; i++  {
		go getJob2Res(jobChan, resultChan)
		//go func() {
		//
		//}()
	}

	for res := range resultChan{
		fmt.Printf("%s result: %d, sum: %d\n",time.Now().Second(), res.job.jobNum, res.sum)
	}

}