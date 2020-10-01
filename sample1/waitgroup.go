package main

import (
	"sync"
)

func Worker(in chan int, i int, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	for {
		println("i=", i)
		num, ok := <-in
		if !ok {
			return
		}
		println("num=", num)
		if num%2 == 0 {
			continue
		}
	}
}

func Dispatcher(out chan int) (err error) {

	defer close(out)

	for i := 0; i < 1000; i++ {
		out <- i
	}
	return
}

func main() {
	chanLimit := 100

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(chanLimit)
	for i := 0; i < chanLimit; i++ {
		go Worker(ch, i, &wg)
	}

	Dispatcher(ch)

	wg.Wait()
	return
}
