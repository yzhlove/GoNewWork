package cnc

import (
	"errors"
	"net"
	"sync"
)

var storage sync.Map

type record struct {
	cc  *net.UDPConn
	err error
}

func take(address string) *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			addr, err := net.ResolveUDPAddr("udp", address)
			if err != nil {
				return &record{err: err}
			}
			cc, err := net.DialUDP("udp", nil, addr)
			return &record{cc: cc, err: err}
		},
	}
}

func Report(address, data string) error {

	var pool *sync.Pool

	res, ok := storage.Load(address)
	if !ok {
		pool = take(address)
		storage.Store(address, pool)
	} else {
		pool = res.(*sync.Pool)
	}

	if rec, ok := pool.Get().(*record); ok {
		//用完将资源归还
		defer pool.Put(rec)
		if rec.err != nil {
			return rec.err
		}
		if _, err := rec.cc.Write([]byte(data)); err != nil {
			return err
		}
		return nil
	}
	return errors.New("type assert error")
}
