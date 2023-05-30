package main

import (
	"fmt"
	"sync"
)

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int, 0)
	wg := sync.WaitGroup{}
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	chA := make(chan int, 5)

	b := [5]int{11, 12, 13, 14, 15}
	chB := make(chan int, 5)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range a {
			chA <- v
		}
		close(chA)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range b {
			chB <- v
		}
		close(chB)
	}()

	chOut := merge(chA, chB)
	for v := range chOut {
		fmt.Println(v)
	}
	wg.Wait()
}
