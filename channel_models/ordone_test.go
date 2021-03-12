package chan_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)


// channel的五种编排模式
// Or-Done模式
// 扇入模式
// 扇出模式
// Stream
// MapReduce

func TestOrDone(t *testing.T) {
	// Or-Done 模式
	// 使用信号通知，实现某个任务执行完成后的通知机制
	// 在实现时，为这个任务定义一个类型为chan struct类型的done变量，
	// 等任务结束后，就可以close这个变量，然后，其他receiver就会收到这个通知
	// 如果有多个任务，只要有任意一个任务执行完成，我们就想获得这个信号，这就是Or-Done模式

	start := time.Now()
	<- or(
		)
	fmt.Printf("Done after %v", time.Since(start))
}

func sig(after time.Duration) <-chan interface{}  {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func or(channels ...<-chan interface{}) <-chan interface{}{

	// 特殊情况下，只有零个或者1和chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			default: // 超过两个，二分法递归处理
				m := len(channels) / 2
				select {
				case <-or(channels[:m]...):
				case <-or(channels[m:]...):
				}
			}
		}
	}()
	return orDone
}

func orReflect(chans ...<-chan interface{}) <-chan interface{}{
	// 利用反射实现Or-Done
	switch len(chans) {
	case 0:
		return nil
	case 1:
		return chans[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		var cases []reflect.SelectCase
		for _, c := range chans{
			cases = append(cases, reflect.SelectCase{
				Dir: reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		// 随机选择一个可用的case
		reflect.Select(cases)
	}()
	return orDone
	
}


