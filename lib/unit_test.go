package lib

import "testing"

func testArrayDataHandle(t *testing.T) {
	d := &Data{}
	d.Init()
	i := 1
	d.ArrayDataHandle(i)
	if len(d.Array) != 1 {
		t.Fail()
	}
}

func testchannelDataHandle(t *testing.T) {
	d := &Data{}
	d.Init()
	i := 1
	d.ChannelDataHandle(i)
	if len(d.Channel) != 1 {
		t.Fail()
	}
}
