package main

import (
	"fmt"
	"time"
)

type statefulWorker struct {
	count int
	ch    chan int
}

func (w *statefulWorker) start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("Current count:", w.count)
			}
		}
	}()
}

func (w statefulWorker) send(value int) {
	w.ch <- value
}

func main() {
	stWorker := &statefulWorker{
		ch: make(chan int),
	}

	stWorker.start()

	for i := range 5 {
		stWorker.send(i)
		time.Sleep(500 * time.Millisecond)
	}
}
