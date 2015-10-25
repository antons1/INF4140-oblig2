package main
import(
	
)

type semaphore chan int;

// The Monitor struct holds information about monitors
type Monitor struct{
	queue semaphore;
	checkin semaphore;
	boarding semaphore;
	riding semaphore;
	unloading semaphore;
}

// ConstructMonitor takes a car capacity as argument, and
// returns a new monitor
func ConstructMonitor(carQSz int) Monitor {
	PrintMessage("Calibrating the monitor...\n");
	var mon Monitor;
	
	mon.queue = make(semaphore);
	mon.checkin = make(semaphore, 1);
	mon.boarding = make(semaphore);
	mon.riding = make(semaphore, 1);
	mon.unloading = make(semaphore);
	
	return mon;
}

// RunMonitor is used to run the Monitor as a goroutine
func (monitor* Monitor) RunMonitor() {
	// Open for passengers
	for {
		
	}
}

// TakeRide is used by Passengers to register themselves into
// the car queue in the monitor. It takes a channel as argument,
// that the monitor uses to notify the car when it has space
func (monitor* Monitor) TakeRide() {
	
	
}

// Load is used by the car to load a passenger
func (monitor* Monitor) Load() {
	
}

// Unload is used by the car to unload a Passenger
func (monitor* Monitor) Unload() {
	
}