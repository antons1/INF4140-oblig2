package main
import(
	"fmt"
)

// The Car struct contains information about a car
type Car struct {
	
}

// ConstructCar creates a new Car and returns it
func ConstructCar() Car {
	ClearMessage();
	fmt.Print("Assembling the car...");
	fmt.Println();
	return Car{};
}

// RunCar is used to run a car as a goroutine
func (car* Car) RunCar(monitor* Monitor) {
	
}