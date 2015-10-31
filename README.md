#Oblig 2 INF4140 - The Roller Coaster problem
An implementation of The Roller Coaster problem of concurrent
programming. Implemented using Go and semaphores.

###How to run
1. Place the source files in a folder called "RollerCoaster" under
your src folder in the GOPATH.
2. Run the command "go install rollercoaster"
3. Find RollerCoaster.exe (RollerCoaster on linux) in your bin
folder in the GOPATH.
4. Run it, and watch 15 digital people run a digital rollercoaster
in a concurrent way!

###Configuration
Do you want to configure the program you say? That's kind of possible!

The following can be edited in rollercoaster.go
```go
const queueSize // The amount of digital people
const carCap	// Spaces in the car
```

The names used for the rides and the passengers can be edited in 
names.go

###Files
| Filename         | File function                      |
| :--------------- | :--------------------------------- |
| rollercoaster.go | Contains main method and spawns the goroutines for the other parts, as well as some utils |
| monitor.go       | Contains most of the program logic, the methods used by car and passenger to synchronize etc. |
| car.go           | Contains the cars logic and data (constructor, wait times etc.) |
| passenger.go     | Contains the passenger logic, constructor, wait times, names etc. |
| semaphore.go     | Contains the semaphore definitions and wait() and signal() functions |
| names.go         | Contains simple random name generator functionality |