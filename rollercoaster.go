package main

// Package rollercoaster implements a solution to solve
// the rollercoaster problem.

import (
	"fmt"
	"math/rand"
	"time"
)

const CarCap = 6	// Max car capacity
const QueueSize = 15	// Number of passengers that are created

var (
	mon  Monitor              // Monitor
	car  Car                  // Car
	pass [QueueSize]Passenger // Passengers
	r1	rand.Rand	// Random Number Generator
)

// Setup initializes everything needed to run the rollercoaster
func setup() {
	fmt.Printf("Building the rollercoaster...\n")
	PrepareRandom();
	mon = ConstructMonitor(CarCap)
	car = ConstructCar(GetCoasterName(), 1)

	for i := range pass {
		pass[i] = ConstructPassenger(GetPersonName(), i)
	}
}

// Main is the main method of RollerCoaster
func main() {
	setup()
	fmt.Println()
	go mon.RunMonitor()
	go car.RunCar(&mon)
	for i := range pass {
		go pass[i].RunPassenger(&mon)
	}

	fmt.Printf("Entering the loops...\n")
	for {

	}
}

// PrepareRandom prepares the RNG
func PrepareRandom() {
	s1 := rand.NewSource(time.Now().UnixNano());
	r1 = *rand.New(s1);
}

// GetRandomNumber generates a random number between min and max
func GetRandomNumber(min, max int) int {
	return r1.Intn(max-min)+min;
}