package factory

import (
	"fmt"
	"valverdethiago/car-factory-threads/assemblyspot"
	"valverdethiago/car-factory-threads/vehicle"
)

const assemblySpots int = 5

type Factory struct {
	vehicleChannel  chan *vehicle.Car
	AssemblingSpots chan *assemblyspot.AssemblySpot
}

func New() *Factory {
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
	}
	return factory
}

// HINT: this function is currently not returning anything, make it concurrent and return right away every single vehicle once assembled,
// (Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(amountOfVehicles int) {
	vehicleList := f.generateVehicleLots(amountOfVehicles)
	vehicleChannel := make(chan *vehicle.Car)

	for i := 1; i <= assemblySpots; i++ {
		go f.worker(i, vehicleChannel)
	}

	for _, vehicle := range vehicleList {
		fmt.Println("Sending vehicle to assembly spot ...")
		vehicleChannel <- vehicle
	}
}

func (Factory) generateVehicleLots(amountOfVehicles int) []*vehicle.Car {
	var vehicles = []*vehicle.Car{}
	var index = 0

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

func (f *Factory) worker(id int, vehicleChannel chan *vehicle.Car) {
	for vehicle := range vehicleChannel {
		fmt.Println("Worker [", id, "] Assembling vehicle...")
		idleSpot := &assemblyspot.AssemblySpot{}
		idleSpot.SetVehicle(vehicle)
		vehicle, err := idleSpot.AssembleVehicle()

		if err != nil {
			fmt.Println("Worker [", id, "] Produced an error while assembling vehicle: ", err.Error())
			continue
		}

		vehicle.TestingLog = f.testCar(vehicle)
		vehicle.AssembleLog = idleSpot.GetAssembledLogs()
	}
}
