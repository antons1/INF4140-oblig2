package main

import (
	"fmt"
)

// The Monitor struct holds information about monitors
type Monitor struct {
	queue     semaphore
	checkin   semaphore
	boarding  semaphore
	riding    semaphore
	unloading semaphore
	inCar     int
}

// ConstructMonitor takes a car capacity as argument, and
// returns a new monitor
func ConstructMonitor(carQSz int) Monitor {
	fmt.Printf("Calibrating the monitor...\n")
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

// TakeRide is used by Passengers to register themselves into
// the car queue in the monitor. It takes a channel as argument,
// that the monitor uses to notify the car when it has space
func (monitor *Monitor) TakeRide() {
	monitor.queue.Wait()
	monitor.checkin.Wait()
	monitor.inCar++
	if monitor.inCar == CarCap {
		monitor.boarding.Signal()
	}
	monitor.checkin.Signal()
	monitor.riding.Wait()
	monitor.unloading.Signal()
}

// Load is used by the car to load a passenger
func (monitor *Monitor) Load() {
	for i := 0; i < CarCap; i++ {
		monitor.queue.Signal()
	}

	monitor.boarding.Wait()
}

// Unload is used by the car to unload a Passenger
func (monitor *Monitor) Unload() {
	monitor.inCar = 0
	for i := 0; i < CarCap; i++ {
		monitor.riding.Signal()
		monitor.unloading.Wait()
	}
}
