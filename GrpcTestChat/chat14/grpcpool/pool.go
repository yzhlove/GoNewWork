package grpcpool

import (
	"errors"
	"google.golang.org/grpc"
	"sync"
)

var (
	errNotConn = errors.New("service busy")
)

type CC struct {
	Conn *grpc.ClientConn
}

type Queue struct {
	host     string
	chConn   chan *CC
	size     int
	capacity int
	sync.RWMutex
}

func New(host string, size, capacity int) *Queue {
	return &Queue{
		host:     host,
		chConn:   make(chan *CC, size),
		size:     size,
		capacity: capacity,
	}
}

func (q *Queue) Get() (*CC, error) {
	select {
	case cc := <-q.chConn:
		return cc, nil
	default:
		q.RLock()
		defer q.RUnlock()
		if len(q.chConn) < q.capacity {
			cc, err := grpc.Dial(q.host, grpc.WithInsecure())
			if err != nil {
				return nil, err
			}
			return &CC{Conn: cc}, nil
		}
		return nil, errNotConn
	}
}

func (q *Queue) Put(cc *CC) {
	q.RLock()
	defer q.RUnlock()

	if len(q.chConn) < q.size {
		select {
		case q.chConn <- cc:
		default:
			return
		}
	}
}
