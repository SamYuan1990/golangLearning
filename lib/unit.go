package lib

import "sync"

type data struct {
	array   []int
	channel chan int
	mutex   sync.Mutex
}

func (d *data) init() {
	d.mutex.Lock()
	d.channel = make(chan int)
	d.mutex.Unlock()
}

func (d *data) arrayDataHandle(i int) {
	d.mutex.Lock()
	d.array = append(d.array, i)
	d.mutex.Unlock()
}

func (d *data) channelDataHandle(i int) {
	d.channel <- i
}
