package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"math/rand"
	"testing"
	pb "think-library/ProtoBuffers/chat08/proto"
	"time"
)

func Test_Memory(t *testing.T) {

	var id uint32 = 1000
	var qty uint32 = 12323453

	rand.Seed(time.Now().UnixNano())

	var items = &pb.Items{} //proto
	var bitems pb.CItems    //tuple
	for i := 0; i < 20; i++ {
		a, b := uint32(rand.Int31()), uint32(rand.Int31())
		fmt.Println(a, "-", b)
		items.List = append(items.List, &pb.Item{Id: a, Qty: b})
		bitems.List = append(bitems.List, pb.CItem{Id: a, Qty: b})
	}

	aaa, _ := proto.Marshal(items)
	bbb, _ := bitems.MarshalMsg(nil)

	fmt.Println("aaaa => ", len(aaa), " bbbb => ", len(bbb))

	a := &pb.Item{Id: id, Qty: qty}
	adata, _ := proto.Marshal(a)

	b := &pb.AItem{Id: id, Qty: qty}
	bdata, _ := b.MarshalMsg(nil)

	c := &pb.BItem{Id: id, Qty: qty}
	cdata, _ := c.MarshalMsg(nil)

	d := &pb.CItem{Id: id, Qty: qty}
	ddata, _ := d.MarshalMsg(nil)

	t.Logf("a:%d b:%d c:%d d:%d", len(adata), len(bdata), len(cdata), len(ddata))
	t.Logf("a:%x b:%x c:%x d:%x", adata, bdata, cdata, ddata)
}

/*
=== RUN   Test_Memory
    main_test.go:26: a:8 b:18 c:15 d:11
--- PASS: Test_Memory (0.00s)
PASS
*/

func Benchmark_Performance(t *testing.B) {

	t.Run("Proto", func(t *testing.B) {
		t.ResetTimer()
		t.StartTimer()
		var items pb.Items
		for i := 0; i < t.N; i++ {
			items.List = append(items.List, &pb.Item{Id: uint32(i), Qty: uint32(i + 1)})
		}

		t.StopTimer()
		t.ReportAllocs()
	})

	t.Run("Msg_A", func(t *testing.B) {
		t.ResetTimer()
		t.StartTimer()

		var items pb.AItems
		for i := 0; i < t.N; i++ {
			items.List = append(items.List, pb.AItem{Id: uint32(i), Qty: uint32(i + 1)})
		}

		t.StopTimer()
		t.ReportAllocs()
	})

	t.Run("Msg_B", func(t *testing.B) {
		t.ResetTimer()
		t.StartTimer()
		var items pb.BItems
		for i := 0; i < t.N; i++ {
			items.List = append(items.List, pb.BItem{Id: uint32(i), Qty: uint32(i + 1)})
		}
		t.StopTimer()
		t.ReportAllocs()
	})

	t.Run("Msg_C", func(t *testing.B) {
		t.ResetTimer()
		t.StartTimer()
		var items pb.CItems
		for i := 0; i < t.N; i++ {
			items.List = append(items.List, pb.CItem{Id: uint32(i), Qty: uint32(i + 1)})
		}
		t.StopTimer()
		t.ReportAllocs()
	})

}

/*
Benchmark_Performance
Benchmark_Performance/Proto
Benchmark_Performance/Proto-12         	20731552	        57.71 ns/op	      57 B/op	       1 allocs/op
Benchmark_Performance/Msg_A
Benchmark_Performance/Msg_A-12         	100000000	        18.62 ns/op	      49 B/op	       0 allocs/op
Benchmark_Performance/Msg_B
Benchmark_Performance/Msg_B-12         	219759943	        17.57 ns/op	      43 B/op	       0 allocs/op
Benchmark_Performance/Msg_C
Benchmark_Performance/Msg_C-12         	172793740	         7.633 ns/op	      44 B/op	       0 allocs/op
PASS
*/
