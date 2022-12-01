package factory

import (
	"fmt"
	"sync"
	"sync/atomic"
	"valverdethiago/car-factory-threads/assemblyspot"
	"valverdethiago/car-factory-threads/vehicle"
)

const assemblySpots int = 5

type Factory struct {
	counter int32
}

func New() *Factory {
	factory := &Factory{
		counter: 0,
	}
	return factory
}

// HINT: this function is currently not returning anything, make it concurrent and return right away every single vehicle once assembled,
// (Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(amountOfVehicles int) {
	vehicleList := f.generateVehicleLots(amountOfVehicles)
	wg := &sync.WaitGroup{}
	wg.Add(amountOfVehicles)
	vehicleChannel := make(chan *vehicle.Car)
	readyVehicleChannel := make(chan *vehicle.Car)

	for i := 1; i <= assemblySpots; i++ {
		go f.worker(i, vehicleChannel, readyVehicleChannel, wg)
	}

	go f.processResults(readyVehicleChannel, amountOfVehicles)

	for _, vehicle := range vehicleList {
		fmt.Println("Sending vehicle to assembly spot ...")
		vehicleChannel <- vehicle
	}
	wg.Wait()
}

func (*Factory) generateVehicleLots(amountOfVehicles int) []*vehicle.Car {
	var vehicles []*vehicle.Car
	var index int
	index = 0

	for {
		vehicles = append(vehicles, &vehicle.Car{
			Id:            index,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
		})

		index++

		if index >= amountOfVehicles {
			break
		}
	}

	return vehicles
}

func (f *Factory) testCar(car *vehicle.Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}

func (f *Factory) worker(id int, vehicleChannel chan *vehicle.Car, readyVehicleChannel chan *vehicle.Car, wg *sync.WaitGroup) {
	for vehicle := range vehicleChannel {
		fmt.Println("Worker [", id, "] Assembling vehicle...")
		idleSpot := &assemblyspot.AssemblySpot{}
		idleSpot.SetVehicle(vehicle)
		f.assembleVehicle(id, idleSpot, readyVehicleChannel, wg)
	}
}

func (f *Factory) assembleVehicle(id int, idleSpot *assemblyspot.AssemblySpot, readyVehicleChannel chan *vehicle.Car, wg *sync.WaitGroup) {
	defer wg.Done()
	vehicle, err := idleSpot.AssembleVehicle()

	if err != nil {
		fmt.Println("Worker [", id, "] Produced an error while assembling vehicle: ", err.Error())
	}

	vehicle.TestingLog = f.testCar(vehicle)
	vehicle.AssembleLog = idleSpot.GetAssembledLogs()
	readyVehicleChannel <- vehicle
}

func (f *Factory) processResults(readyVehicleChannel chan *vehicle.Car, amountOfVehicles int) {
	for ready := range readyVehicleChannel {
		atomic.AddInt32(&f.counter, 1)
		fmt.Println("Vehicle ", f.counter, " of ", amountOfVehicles, " assembly process is finished")
		fmt.Println("\tAssembly logs: ", ready.AssembleLog)
		fmt.Println("\tTesting logs: ", ready.TestingLog)
	}
}
