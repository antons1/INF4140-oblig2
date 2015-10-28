package main

import (
	"time"
)

// Define a semaphore
type semaphore chan int

// Wait sets the semaphore to wait
func (s semaphore) Wait() {
	time.Sleep(200 * time.Millisecond) // Slow down to make output readable
	s <- 1
}

// Signal sets the semaphore to continue
func (s semaphore) Signal() {
	time.Sleep(200 * time.Millisecond) // Slow down to make output readable
	<-s
}
