package lib

import "testing"

func testArrayDataHandle(t *testing.T) {
	d := &data{}
	d.init()
	i := 1
	d.arrayDataHandle(i)
	if len(d.array) != 1 {
		t.Fail()
	}
}

func testchannelDataHandle(t *testing.T) {
	d := &data{}
	d.init()
	i := 1
	d.channelDataHandle(i)
	if len(d.channel) != 1 {
		t.Fail()
	}
}
