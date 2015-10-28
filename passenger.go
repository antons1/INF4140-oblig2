package main

import (
	"fmt"
	"time"
)

// The Passenger struct contains information about a passenger
type Passenger struct {
	name string	// For terminal output
	id   int	// To discern between passengers with the same name
}

// ConstructPassenger creates a new instantiated passenger
// struct and returns it
func ConstructPassenger(nm string, id int) Passenger {
	fmt.Printf("## Straightening passenger spines (%d passengers)...\r", id)
	return Passenger{nm, id}
}

// RunPassenger is used to run a passenger as a goroutine and print output
func (passenger *Passenger) RunPassenger(monitor *Monitor) {
	for {
		fmt.Printf("## %-8s ## %s (%d) queuing up for ride\n",
			"PASS",
			passenger.name,
			passenger.id)

		// Added to the queue
		monitor.TakeRide()

		fmt.Printf("## %-8s ## %s (%d) got off the ride and is wandering around the park\n",
			"PASS",
			passenger.name,
			passenger.id)

		time.Sleep(time.Duration(GetRandomNumber(5, 15)) * time.Second)
	}
}
