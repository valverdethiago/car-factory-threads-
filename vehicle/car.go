package vehicle

import "fmt"

type Car struct {
	Id            int
	Chassis       string
	Tires         string
	Engine        string
	Electronics   string
	Dash          string
	Sits          string
	Windows       string
	EngineStarted bool
	TestingLog    string
	AssembleLog   string
}

func (c *Car) StartEngine() (string, error) {
	if c.EngineStarted {
		return "", fmt.Errorf("Cannot start engine already started")
	}

	return "Engine Started!", nil
}

func (c *Car) StopEngine() (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("Cannot stop engine already stopped")
	}

	return "Engine Stopped!", nil
}

func (c *Car) MoveForwards(distance int) (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("Cannot move with stopped engine")
	}

	return fmt.Sprintf("Moved forward %d meters!", distance), nil
}

func (c *Car) MoveBackwards(distance int) (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("Cannot move with stopped engine")
	}

	return fmt.Sprintf("Moved backwards %d meters!", distance), nil
}

func (c *Car) TurnRight() (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("Cannot turn right with stopped engine")
	}

	return fmt.Sprintf("Turned Right!"), nil
}

func (c *Car) TurnLeft() (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("Cannot turn left with stopped engine")
	}

	return fmt.Sprintf("Turned Right"), nil
}
