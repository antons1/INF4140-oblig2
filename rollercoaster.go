package main
// Package rollercoaster implements a solution to solve
// the rollercoaster problem.

import(
	"fmt"
)

const CarCap = 6; // Max car capacity
const QuSize = 10; // Size of passenger queue

var (
	mon Monitor; // Monitor
	car Car; // Car
	pass[QuSize] Passenger; // Passengers
)

// Setup initializes everything needed to run the rollercoaster
func setup() {
	fmt.Print("Initializing the rollercoaster...");
	fmt.Println();
	mon = ConstructMonitor(CarCap);
	car = ConstructCar();
	
	for i := range pass {
		pass[i] = ConstructPassenger();
	}
}

// Main is the main method of RollerCoaster
func main() {
	setup();
	ClearMessage();
	mon.RunMonitor();
	car.RunCar(&mon);
	for i := range pass {
		pass[i].RunPassenger(&mon);
	}
}

// ClearMessage clears the current line
func ClearMessage() {
	fmt.Print("\r                                                  \r");
}