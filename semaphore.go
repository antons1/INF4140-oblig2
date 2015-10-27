package main

import (
	"time"
)

type semaphore chan int

func (s semaphore) Wait() {
	time.sleep(100 * time.Millisecond)
	s <- 1
}

func (s semaphore) Signal() {
	time.sleep(100 * time.Millisecond)
	<-s
}
