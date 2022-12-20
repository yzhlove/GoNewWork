package pipeline

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

func ReaderSource(reader io.Reader) <-chan int {
	var out = make(chan int)
	var buffer = make([]byte, 8)
	go func() {
		for {
			n, err := reader.Read(buffer)
			if err != nil {
				break
			}
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(buffer))
			}
		}
		close(out)
	}()
	return out
}

func WriterSink(w io.Writer, in <-chan int) error {
	for x := range in {
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, uint64(x))
		if _, err := w.Write(bytes); err != nil {
			return err
		}
	}
	return nil
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- int(rand.Int31n(100))
		}
		close(out)
	}()
	return out
}

func MemorySource(numbers <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		cache := make([]int, 0, 128)
		for x := range numbers {
			cache = append(cache, x)
		}
		sort.Ints(cache)
		for _, x := range cache {
			out <- x
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		x1, ok1 := <-in1
		x2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && x1 <= x2) {
				out <- x1
				x1, ok1 = <-in1
			} else {
				out <- x2
				x2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}
