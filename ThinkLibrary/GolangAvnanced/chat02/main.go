package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//test1()
	//test2()

	a := make([]int, 4, 16)
	fmt.Println(a[4:])

}

func test1() {

	// 限制当前只有p
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		i := i
		// 在当前的p里面创建10个go携程（下文的G），
		// 每个G都会想到runnext，在进入本地的local queue.
		// 因为runext只能容纳一个G
		// 所以只有最后一个G进入了runnext
		// 其余的G都按顺序进入了local queue
		go func() {
			fmt.Println("A:", i)
		}()
	}

	ch := make(chan struct{})
	<-ch
}

/*
A: 9
A: 0
A: 1
A: 2
A: 3
A: 4
A: 5
A: 6
A: 7
A: 8
fatal error: all goroutines are asleep - deadlock!
*/

func test2() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println("A:", i)
		}()
	}

	// 在go1.14之前time.sleep会创建一个G,
	// time.Sleep里面的这个G是最后一个，因而会进入runnext
	// 所以在go1.14之前输出的内容为 0 1 2 3 4 5 6 7 8 9
	// 在go1.14之后因为time.sleep不会创建G了，所以输出同上
	time.Sleep(time.Minute)
}

/*
A: 9
A: 0
A: 1
A: 2
A: 3
A: 4
A: 5
A: 6
A: 7
A: 8
*/
