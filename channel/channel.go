package channel

import (
	"bitbucket.org/biglog/writeDisk"
	"math/rand"
	"runtime"
	"time"
)

//随机模式，其实也可以用原子加和swap的方法进行轮询
//TODO:验证原子操作好还是随机模式好
func SendLog(log []byte) {
	rand.Seed(time.Now().UnixNano())
	channelIndex := rand.Int() % runtime.NumCPU()

	channelSlice[channelIndex] <- log
}

var channelSlice []chan []byte

//创建CPU数量的协程来将日志写入磁盘
func init() {
	channelSlice = make([]chan []byte, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		t := make(chan []byte, 1024)
		channelSlice[i] = t

		go func() {
			writeDisck := writeDisk.NewWriteDisck()
			for {
				log := <-t
				writeDisck.WriteLog(log)
			}
		}()
	}
}
