package main

import (
	"fmt"
	"time"
)

// The Car struct contains information about a car
type Car struct {
	name string
	id   int
}

// ConstructCar creates a new Car and returns it
func ConstructCar(nm string, id int) Car {
	fmt.Printf("## Assembling the car...\n")
	return Car{nm, id}
}

// RunCar is used to run a car as a goroutine
func (car *Car) RunCar(monitor *Monitor) {
	for {
		fmt.Printf("## %-8s ## %s is ready\n",
			"CAR",
			car.name)
		fmt.Printf("## %-8s ## %s is loading passengers\n",
			"CAR",
			car.name)
		monitor.Load()

		fmt.Printf("## %-8s ## %s rides around the track\n",
			"CAR",
			car.name)
		time.Sleep(5 * time.Second)

		fmt.Printf("## %-8s ## %s is unloading passengers\n",
			"CAR",
			car.name)
		monitor.Unload()
	}
}
