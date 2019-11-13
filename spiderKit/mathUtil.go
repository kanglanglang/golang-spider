package spiderKit

import (
	"math/rand"
	"sync"
	"time"
)

var (
	// 等待组
	downloadWG sync.WaitGroup
	// 互斥锁
	randomMT sync.Mutex
)

// GetRandomInt 获取随机数 用互斥锁锁防止并发同名
func GetRandomInt(start, end int) int {
	randomMT.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := start + r.Intn(end-start)
	randomMT.Unlock()
	return ret
}
