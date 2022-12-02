package geerpc

import (
	"7days-golang/gee-rpc/day1-codec/geerpc/codec"
	"sync"
)

type Call struct {
	Seq           uint64
	ServiceMethod string
	Args          interface{}
	Reply         interface{}
	Error         error
	Done          chan *Call
}

type Client struct {
	cc       codec.Codec
	opt      Option
	sending  sync.Mutex
	header   codec.Header
	mu       sync.Mutex
	seq      uint64
	pending  map[uint64]*Call
	closing  bool
	shutdown bool
}

func (call *Call) done() {
	call.Done <- call
}
