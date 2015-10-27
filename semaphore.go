package main

import (
	"time"
)

// Define a semaphore
type semaphore chan int

// Wait sets the semaphore to wait
func (s semaphore) Wait() {
	time.Sleep(100 * time.Millisecond)
	s <- 1
}

// Signal sets the semaphore to continue
func (s semaphore) Signal() {
	time.Sleep(100 * time.Millisecond)
	<-s
}
