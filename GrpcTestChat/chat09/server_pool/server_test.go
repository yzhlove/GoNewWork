package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test-chat/chat09/echo"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

type Conn struct {
	conn echo.EchoServiceClient
	err  error
}

type Queue struct {
	cp *sync.Pool
}

func NewQueue() *Queue {
	return &Queue{
		cp: &sync.Pool{
			New: func() interface{} {
				cc, err := grpc.Dial(_port, grpc.WithInsecure())
				if err != nil {
					return &Conn{err: err}
				}
				c := echo.NewEchoServiceClient(cc)
				return &Conn{conn: c, err: nil}
			},
		},
	}

}

func (q *Queue) Get() *Conn {
	return q.cp.Get().(*Conn)
}

func (q *Queue) Put(cc *Conn) {
	q.cp.Put(cc)
}

func Test_Grpc(t *testing.T) {
	queue := NewQueue()
	rand.Seed(time.Now().UnixNano())

	count := 10000
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			cc := queue.Get()
			if cc.err != nil {
				panic(fmt.Sprintf("get connection error: %v", cc.err))
			}
			defer queue.Put(cc)
			strs, value := generator(rand.Intn(50))
			req := strings.Join(strs, "-")
			resp, err := cc.conn.Echo(context.Background(), &echo.Req{Msg: req})
			if err != nil {
				t.Error(err)
				return
			}

			if fmt.Sprintf("%d", value) != resp.Msg {
				t.Error(fmt.Sprintf("cacel result noinstanic, value:%d ,real:%s \n", value, resp.Msg))
				return
			}

		}(i)
	}

	wg.Wait()

}

func generator(size int) ([]string, int) {
	var sum int
	var strs = make([]string, 0, size)
	for i := 0; i < size; i++ {
		number := rand.Intn(1000)
		sum += number
		strs = append(strs, strconv.Itoa(number))
	}

	return strs, sum
}
