package main
import(
	"fmt"
)

// The Monitor struct holds information about monitors
type Monitor struct{
	trOpen chan int // 
	crChans[] chan int
}

// ConstructMonitor takes a car capacity as argument, and
// returns a new monitor
func ConstructMonitor(carQSz int) Monitor {
	var mon Monitor;
	mon.trOpen = make(chan int);
	mon.crChans = make([]chan int, 0, 1);
	
	return mon;
}

// RunMonitor is used to run the Monitor as a goroutine
func (monitor* Monitor) RunMonitor() {
	for {
		
	}
}

// TakeRide is used by Passengers to register themselves into
// the car queue in the monitor. It takes a channel as argument,
// that the monitor uses to notify the car when it has space
func (monitor* Monitor) TakeRide(carReady chan int) {
	monitor.crChans = append(monitor.crChans, carReady);
}

// Load is used by the car to load a passenger
func (monitor* Monitor) Load() {
	
}

// Unload is used by the car to unload a Passenger
func (monitor* Monitor) Unload() {
	
}