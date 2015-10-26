package main

import ()

type semaphore chan int

func (s semaphore) Wait() {
	s <- 1
}

func (s semaphore) Signal() {
	<-s
}
