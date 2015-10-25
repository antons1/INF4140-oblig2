package main
import(
	"fmt"
)

// The Passenger struct contains information about a passenger
type Passenger struct {
	
}

// ConstructPassenger creates a new instantiated passenger
// struct and returns it
func ConstructPassenger() Passenger {
	ClearMessage();
	fmt.Print("Straightening passenger spines...");
	return Passenger{};
}

// RunPassenger is used to run a passenger as a goroutine
func (passenger* Passenger) RunPassenger(monitor* Monitor) {
	
}