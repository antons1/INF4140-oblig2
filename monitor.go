package main

import (
	"fmt"
)

// The Monitor struct holds information about monitors
type Monitor struct {
	// Wait/continue semaphores
	queue     semaphore // Passengers queue up (no limit)
	checkin   semaphore // Passengers check in to the ride (1 by 1)
	boarding  semaphore // Passengers board the car (1 by 1)
	riding    semaphore // Car is riding
	unloading semaphore // Passengers are unloading (1 by 1)
	inCar     int       // How many passengers are in the car
}

// ConstructMonitor takes a car capacity as argument, and
// returns a new monitor
func ConstructMonitor(carQSz int) Monitor {
	fmt.Printf("## Calibrating the monitor...\n")
	var mon Monitor

	mon.queue = make(semaphore, 1)
	mon.checkin = make(semaphore, 1)
	mon.boarding = make(semaphore, 1)
	mon.riding = make(semaphore, 1)
	mon.unloading = make(semaphore, 1)

	mon.queue.Wait()
	mon.boarding.Wait()
	mon.riding.Wait()
	mon.unloading.Wait()

	return mon
}

// RunMonitor is used to run the Monitor as a goroutine
func (monitor *Monitor) RunMonitor() {
	for {

	}
}

// TakeRide is used by Passengers to queue up for the Car.
// The function is halted on the different semaphores, effectively
// making the calling Passenger wait until it is signaled to continue
func (monitor *Monitor) TakeRide() {
	// First barrier, Passengers wait to be let into the checkin area
	// This is so that the car can control that only C people are allowed
	// to continue (Semaphore is set on wait to begin with, so that passengers
	// are not allowed to continue until the car is ready)
	monitor.queue.Wait()

	// Second barrier, after being let through the queue, the first passenger
	// can directly enter the car. The next passenger has to wait for the
	// previous passenger to signal before continuing with checkin.
	// The inCar variable is increased to see how many are boarded
	monitor.checkin.Wait()
	monitor.inCar++
	fmt.Printf("## %-8s ## Now %d passengers in car (boarding)\n",
		"MONITOR",
		monitor.inCar)

	// After C passengers has boarded, the last passenger notifies the car
	// that boarding is complete
	if monitor.inCar == CarCap {
		monitor.boarding.Signal()
	}

	// Here check-in is reopened, to let the next passenger through. If inCar < C,
	// that means the next passenger in the queue. If inCar = C, the car will not
	// have let any more passengers into the check-in area, and check-in will be
	// ready for next trip
	monitor.checkin.Signal()

	// Passenger waits until car signals that riding has stopped. This makes sure that
	// the Passenger stays in the car until it has gone around the track and stopped
	// again
	monitor.riding.Wait()

	// The Passenger signals the car that it has finished unloading, so the car can
	// either let the next passenger out, or start admitting new passengers
	monitor.unloading.Signal()
}

// Load is used by the car to load all passengers
func (monitor *Monitor) Load() {
	// First, let C passengers into the check-in area to make sure that the car
	// can be filled
	for i := 0; i < CarCap; i++ {
		monitor.queue.Signal()
	}

	// Then wait until all passengers have boarded
	monitor.boarding.Wait()
}

// Unload is used by the car to unload al Passengers
func (monitor *Monitor) Unload() {
	// inCar is decreased to 0 again during this process
	for ; 0 < monitor.inCar; monitor.inCar-- {
		// Tell passengers that the ride has stopped
		monitor.riding.Signal()

		// Wait for passengers to unload
		monitor.unloading.Wait()
		fmt.Printf("## %-8s ## Now %d passengers in car (unboarding)\n",
			"MONITOR",
			monitor.inCar)
	}
}
