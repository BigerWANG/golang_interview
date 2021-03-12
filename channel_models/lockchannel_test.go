package chan_test

import (
	"testing"
	"time"
)


type CMutex struct {
	ch chan struct{}
}

// 使用锁需要初始化
func NewCMutex() *CMutex {
	mu := &CMutex{
		make(chan struct{}, 1),
	}
	mu.ch <- struct{}{}
	return mu
}

// 请求锁，直到获取
func(m *CMutex)Lock(){
	<-m.ch
}

// 解锁
func(m *CMutex)Unlock(){
	select {
	case m.ch<- struct{}{}:
	default:
		panic("unlock of unlock mutex")
	}
}

// 尝试获取锁
func (m *CMutex) TryLock() bool{
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}


// 加入一个超时的设置
func (m *CMutex)LockTimeout(timeout time.Duration)bool  {
	timer := time.NewTimer(timeout)
	select {
	case <- timer.C:
	case <-m.ch:
		timer.Stop()
		return true
	}
	return false
}


// 检查锁是否已经被持有
func (m *CMutex)IsLocked() bool {
	return len(m.ch) == 0
}




func TestChanLock(t *testing.T) {
	// 使用chan实现互斥锁
	cmu := NewCMutex()
	cmu.Lock()
	t.Log(cmu.IsLocked())
	cmu.LockTimeout(time.Second * 3)
	t.Log(cmu.TryLock())
	t.Log(cmu.IsLocked())
	cmu.Unlock()
	t.Log(cmu.IsLocked())

}
