package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/SamYuan1990/golangLearning/lib"
)

// wait group & lock
func waitGroup() {
	d := &lib.Data{}
	d.Init()
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(n int) {
			d.ArrayDataHandle(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(d.Array))
}

// pointer into mutiple channel
func pointerIntoMutipleChannel() {
	channel1 := make(chan *lib.Data, 1)
	channel2 := make(chan *lib.Data, 1)
	d := &lib.Data{}
	d.Init()
	channel1 <- d
	channel2 <- d
	d1 := <-channel1
	d2 := <-channel2
	fmt.Println(d1 == d2)
}

// single summary channel
func singleSummaryChannel() {
	//origin channel
	//processor 1 get d from origin channel if not 1, +1 into channel of d
	//processor 2 get d from origin channel if not 2, +2 into channel of d
	//chan check check if d.channel len == 2 else back to origin channel
	originChan := make(chan *lib.Data, 10)
	checkChan := make(chan *lib.Data, 10)
	processor1 := &lib.Processor{No: 1}
	processor2 := &lib.Processor{No: 2}
	done := make(chan struct{})
	go processor1.LoopChannel(originChan, checkChan, done)
	go processor2.LoopChannel(originChan, checkChan, done)
	go processor1.Checking(checkChan, originChan, done)
	go processor2.Checking(checkChan, originChan, done)
	for i := 0; i < 5; i++ {
		data := &lib.Data{}
		data.Init()
		fmt.Println(data)
		originChan <- data
	}
	time.Sleep(time.Duration(1) * time.Second)
	close(done)
	fmt.Println(len(originChan))
	fmt.Println(len(checkChan))
}

// loop channel
/*
for {
					if data, ok := <-r.Channel; ok {
						fmt.Println(r, data)
					} else {
						break
					}
				}
				//close(r.Channel)
*/

// mutiple summary channel with pointer
func mutipleSummaryChannelWithPointer() {
	channel1 := make(chan *lib.Data, 1)
	channel2 := make(chan *lib.Data, 1)
	processor1 := &lib.Processor{No: 1}
	processor2 := &lib.Processor{No: 2}
	done := make(chan struct{})
	checkChan := make(chan *lib.Data, 10)
	go processor1.LoopChannel(channel1, checkChan, done)
	go processor2.LoopChannel(channel2, checkChan, done)
	go processor1.Checking(checkChan, channel1, done)
	go processor2.Checking(checkChan, channel2, done)
	d := &lib.Data{}
	d.Init()
	channel1 <- d
	channel2 <- d
	time.Sleep(time.Duration(1) * time.Second)
	close(done)
	fmt.Println(len(checkChan))
}

func main() {
	waitGroup()
	pointerIntoMutipleChannel()
	singleSummaryChannel()
	mutipleSummaryChannelWithPointer()
}
