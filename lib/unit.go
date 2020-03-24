package lib

import (
	"fmt"
	"sync"
)

type Data struct {
	Array   []int
	Channel chan int
	Flag    int
	Mutex   sync.Mutex
}

func (d *Data) Init() {
	d.Mutex.Lock()
	d.Channel = make(chan int, 10)
	d.Mutex.Unlock()
}

func (d *Data) ArrayDataHandle(i int) {
	d.Mutex.Lock()
	d.Flag = i
	d.Array = append(d.Array, i)
	d.Mutex.Unlock()
}

func (d *Data) ChannelDataHandle(i int) {
	d.Flag = i
	d.Channel <- i
}

type Processor struct {
	No int
}

func (p *Processor) LoopChannel(input, output chan *Data, done <-chan struct{}) {
	for {
		select {
		case r := <-input:
			if r.Flag != p.No {
				r.ArrayDataHandle(p.No)
			}
			output <- r
		case <-done:
			return
		}
	}
}

func (p *Processor) Checking(input, output chan *Data, done <-chan struct{}) {
	for {
		select {
		case r := <-input:
			if len(r.Array) < 2 {
				output <- r
			} else {
				fmt.Println(r, r.Array)
			}
		case <-done:
			return
		}
	}
}
