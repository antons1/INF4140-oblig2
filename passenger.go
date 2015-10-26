package main

import (
	"fmt"
	"time"
)

// The Passenger struct contains information about a passenger
type Passenger struct {
	name string
	id   int
}

// ConstructPassenger creates a new instantiated passenger
// struct and returns it
func ConstructPassenger(nm string, id int) Passenger {
	PrintMessage("Straightening passenger spines...")
	return Passenger{nm, id}
}

// RunPassenger is used to run a passenger as a goroutine
func (passenger *Passenger) RunPassenger(monitor *Monitor) {
	for {
		fmt.Printf("%-10s %s %d queuing up for ride\n",
			"#PASS",
			passenger.name,
			passenger.id)

		// Added to the queue
		monitor.TakeRide()

		fmt.Printf("%-10s %s %d got off the ride and is wandering around the park\n",
			"#PASS",
			passenger.name,
			passenger.id)

		time.Sleep(5 * time.Second)
	}
}
