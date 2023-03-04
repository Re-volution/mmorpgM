package main

import "sync/atomic"

var idddd = new(uint32)

func getid() uint32 {
	return atomic.AddUint32(idddd, 1)
}
