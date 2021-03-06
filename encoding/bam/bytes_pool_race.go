// Code generated by " ../../../base/gtl/generate.py --output=bytes_pool_race.go --prefix=bytes -DELEM=[]byte --package=bam ../../../base/gtl/randomized_freepool_race.go.tpl ". DO NOT EDIT.

// +build race

package bam

import "sync/atomic"

type bytesFreePool struct {
	new func() []byte
	len int64
}

func NewbytesFreePool(new func() []byte, maxSize int) *bytesFreePool {
	return &bytesFreePool{new: new}
}

func (p *bytesFreePool) Put(x []byte) {
	atomic.AddInt64(&p.len, -1)
}

func (p *bytesFreePool) Get() []byte {
	atomic.AddInt64(&p.len, 1)
	return p.new()
}

func (p *bytesFreePool) ApproxLen() int { return int(atomic.LoadInt64(&p.len)) }
