package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

var (
	endpoints = []string{"localhost:2379"}
	timeout   = 5 * time.Second
)

const (
	KeyCreate = iota
	KeyUpdate
	KeyDelete
)

type Etcd struct {
	endpoints []string
	cc        *clientv3.Client
	kv        clientv3.KV
	timeout   time.Duration
}

type KeyChangedEvent struct {
	Type  int
	Key   string
	Value []byte
}

type WatcherChangedResponse struct {
	Event      chan *KeyChangedEvent
	CancelFunc context.CancelFunc
	Watcher    clientv3.Watcher
}

type Response struct {
	Key, Value string
}

type TxResponse struct {
	Succeed bool
	LeaseID clientv3.LeaseID
	Lease   clientv3.Lease
	Key     string
	Value   string
}

func NewEtcd(endpoints []string, timeout time.Duration) (etcd *Etcd, err error) {
	var cc *clientv3.Client
	if cc, err = clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: timeout}); err != nil {
		return
	}
	etcd = &Etcd{endpoints: endpoints, cc: cc, kv: clientv3.NewKV(cc), timeout: timeout}
	return
}

func (e *Etcd) Get(key string) (resps []Response, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	resp, err := e.kv.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) > 0 {
		for _, r := range resp.Kvs {
			resps = append(resps, Response{Key: string(r.Key), Value: string(r.Value)})
		}
	}
	return
}

func (e *Etcd) Put(key, value string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	_, err = e.kv.Put(ctx, key, value)
	return
}

func (e *Etcd) Delete(key string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	_, err = e.kv.Delete(ctx, key, clientv3.WithPrefix())
	return
}

func (e *Etcd) handleEvent(evt *clientv3.Event, ch chan *KeyChangedEvent) {
	changedEvent := &KeyChangedEvent{Key: string(evt.Kv.Key)}
	switch evt.Type {
	case clientv3.EventTypePut:
		changedEvent.Type = KeyUpdate
		if evt.IsCreate() {
			changedEvent.Type = KeyCreate
		}
		changedEvent.Value = evt.Kv.Value
	case clientv3.EventTypeDelete:
		changedEvent.Type = KeyDelete
	}
	ch <- changedEvent
}

func (e *Etcd) Watch(key string) *WatcherChangedResponse {

	w := clientv3.NewWatcher(e.cc)
	ch := w.Watch(context.Background(), key, clientv3.WithPrefix())
	watcherResp := &WatcherChangedResponse{
		Event:   make(chan *KeyChangedEvent, 128),
		Watcher: w,
	}

	go func() {
		for rec := range ch {
			if rec.Canceled {
				break
			}
			for _, evt := range rec.Events {
				e.handleEvent(evt, watcherResp.Event)
			}
		}
	}()
	return watcherResp
}

func main() {

	e, err := NewEtcd(endpoints, timeout)
	if err != nil {
		log.Fatal(err)
	}

	resp := e.Watch("/root")
	for evt := range resp.Event {
		fmt.Println("evt value => ", evt)
	}
}
