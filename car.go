package main

import (
	"fmt"
	"time"
)

// The Car struct contains information about a car
type Car struct {
	name string // Used for terminal output
	id   int    // In case of multiple cars in the future
}

// ConstructCar creates a new Car and returns it
func ConstructCar(nm string, id int) Car {
	fmt.Printf("## Assembling the car...\n")
	return Car{nm, id}
}

// RunCar is used to run a car as a goroutine, with terminal output
func (car *Car) RunCar(monitor *Monitor) {
	for {
		fmt.Printf("## %-8s ## %s enters the station\n",
			"CAR",
			car.name)
		fmt.Printf("## %-8s ## %s is loading passengers\n",
			"CAR",
			car.name)
		monitor.Load() // Car waits for passengers to enter

		fmt.Printf("## %-8s ## %s rides around the track\n",
			"CAR",
			car.name)
		time.Sleep(15 * time.Second)

		fmt.Printf("## %-8s ## %s is unloading passengers\n",
			"CAR",
			car.name)
		monitor.Unload() // Car waits for passengers to leave
		fmt.Printf("## %-8s ## %s is empty\n",
			"CAR",
			car.name)

	}
}
