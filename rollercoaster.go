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
	PrintMessage("Building the rollercoaster...\n");
	mon = ConstructMonitor(CarCap);
	car = ConstructCar("Roaring Waterfalls", 1);
	
	for i := range pass {
		pass[i] = ConstructPassenger("Hermann", i);
	}
}

// Main is the main method of RollerCoaster
func main() {
	setup();
	fmt.Println();
	go mon.RunMonitor();
	go car.RunCar(&mon);
	for i := range pass {
		go pass[i].RunPassenger(&mon);
	}
	
	PrintMessage("Entering the loops...\n");
	for {
		
	}
}

// ClearMessage clears the current line
func PrintMessage(arg string) {
	fmt.Printf("%20s\r", arg);
}