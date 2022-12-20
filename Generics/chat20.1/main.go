package main

import (
	"fmt"
	"generics-chat/chat20.1/pipeline"
	"os"
)

func main() {

	file, err := os.Open("test1.binary")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	out := pipeline.MemorySource(pipeline.ReaderSource(file))
	for x := range out {
		fmt.Print(x, " ")
	}
}

func write() {
	file, err := os.Create("test1.binary")
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
