package chan_test

import (
	"reflect"
	"testing"
)

func TestFanIn(t *testing.T) {
	// 在软件工程中，模块的扇入是指有多少个上级模块调用它。
	// channel的扇入模式是指有多个源Channel输入，一个目的channel输出的情况
	//

}

func fanInReflect(chans ...<-chan interface{}) <-chan interface{}{
	out := make(chan interface{})
	go func() {
		defer close(out)
		// 构造SelectCase slice
		var cases []reflect.SelectCase
		for _, c := range chans{
			cases = append(cases, reflect.SelectCase{
				Dir: reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 循环，从cases中选择一个可用的
		for len(cases) > 0{
			i, v, ok := reflect.Select(cases)
			if !ok{ // 此时channel已经close
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()
	return out
}
