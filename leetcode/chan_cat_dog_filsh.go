package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count uint32
	mutex *sync.Mutex
}

func (c *Counter) Add() {
	c.mutex.Lock()
	c.count += 1
	c.mutex.Unlock()
}

func (c *Counter) Get() uint32 {
	var res uint32
	c.mutex.Lock()
	res = c.count
	c.mutex.Unlock()
	return res
}

func main() {
	catCh := make(chan int, 1)
	dogCh := make(chan int, 1)
	fishCh := make(chan int, 1)
	counter := Counter{
		count: uint32(0),
		mutex: &sync.Mutex{},
	}

	catCh <- 1
	var wg sync.WaitGroup
	wg.Add(3)
	go printCat(&wg, counter, catCh, dogCh)
	go printDog(&wg, counter, dogCh, fishCh)
	go printFish(&wg, counter, fishCh, catCh)
	wg.Wait()
}

func printCat(wg *sync.WaitGroup, counter Counter, catCh, dogCh chan int) {
	for {
		if counter.Get() >= 100 {
			wg.Done()
			return
		}
		select {
		case <-catCh:
			fmt.Println("cat")
			counter.Add()
			dogCh <- 1
		}
	}
}

func printDog(wg *sync.WaitGroup, counter Counter, dogCh, fishCh chan int) {
	for {
		if counter.Get() >= 100 {
			wg.Done()
			return
		}
		select {
		case <-dogCh:
			fmt.Println("dog")
			counter.Add()
			fishCh <- 1
		}
	}
}

func printFish(wg *sync.WaitGroup, counter Counter, fishCh, catCh chan int) {
	for {
		if counter.Get() >= 100 {
			wg.Done()
			return
		}
		select {
		case <-fishCh:
			fmt.Println("fish")
			counter.Add()
			catCh <- 1
		}
	}
}
