package main

import (
	"bufio"
	"fmt"
	"generics-chat/chat20.1/pipeline"
	"io"
	"os"
)

func main() {
	entry()
}

func read() {
	file, err := os.Open("test1.binary")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	out := pipeline.MemorySort(pipeline.ReaderSource(file))
	for x := range out {
		fmt.Print(x, " ")
	}
}

func write() {
	file, err := os.Create("test.in")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = pipeline.WriterSink(file, pipeline.RandomSource(50))
	if err != nil {
		fmt.Println("write file error")
		panic(err)
	}
	fmt.Println("write file Ok.")
}

func merge(s ...int) int {
	if len(s) == 1 {
		fmt.Println("return s[0]:", s[0])
		return s[0]
	}
	m := len(s) / 2
	return add(merge(s[:m]...), merge(s[m:]...))
}

func add(a, b int) int {
	fmt.Println(a, b)
	return a + b
}

func entry() {

	//pipe, err := createdPipe("test.in")
	pipe, err := createNetworkPipe("test.in")
	if err != nil {
		panic(err)
	}
	if err := writeToFile(pipe, "test.out"); err != nil {
		panic(err)
	}

	if err := printFile("test.out"); err != nil {
		panic(err)
	}
}

func createdPipe(path string) (<-chan int, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	var count = 6
	var chunk = getChunk(int(stat.Size()), count)
	var sortRes = make([]<-chan int, 0, count)
	for i := 0; i < count; i++ {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		file.Seek(int64(i*chunk), io.SeekStart)
		sortRes = append(sortRes, pipeline.MemorySort(pipeline.ReaderSourceAt(bufio.NewReader(file), chunk)))
	}
	return pipeline.MergeN(sortRes...), nil
}

func createNetworkPipe(path string) (<-chan int, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	var count = 6
	var chunk = getChunk(int(stat.Size()), count)

	var address []string
	for i := 0; i < count; i++ {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		file.Seek(int64(i*chunk), io.SeekStart)

		source := pipeline.ReaderSourceAt(bufio.NewReader(file), chunk)
		addr := fmt.Sprintf(":780%d", i)
		if err := pipeline.NetworkSink(addr, pipeline.MemorySort(source)); err != nil {
			return nil, err
		}

		address = append(address, addr)
	}

	var res []<-chan int
	for _, addr := range address {
		res = append(res, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(res...), nil
}

func writeToFile(in <-chan int, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	return pipeline.WriterSink(writer, in)
}

func printFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var index int
	for x := range pipeline.ReaderSource(bufio.NewReader(file)) {
		index++
		fmt.Println(index, x)
	}
	return nil
}

func getChunk(size, bucket int) int {
	if size%8 != 0 {
		panic("size is error")
	}

	count := size / 8
	if count%bucket != 0 {
		return (count/bucket + 1) * 8
	}
	return count / bucket * 8
}
