package marsrover

import "fmt"

const (
	Xmax = 10
	Ymax = 10
)

type Rover interface {
	getPosition() (int, int)
	move(amount int) Rover
	turnLeft() Rover
	turnRight() Rover
}

type Instructions []func(Rover) (Rover, error)

func Perform(rover Rover, instructions Instructions) (r Rover, e error) {
	r = rover
	for _, instr := range instructions {
		r, e = instr(r)
		if e != nil {
			return
		}
	}
	return
}

func Forward(rover Rover) (Rover, error) {
	moved := rover.move(1)

	if collidesWithObstacle(moved) {
		return rover, fmt.Errorf("obstacle detected")
	}

	return moved, nil
}

func Backward(rover Rover) (Rover, error) {
	return rover.move(-1), nil
}

func Left(rover Rover) (Rover, error) {
	return rover.turnLeft(), nil
}

func Right(rover Rover) (Rover, error) {
	return rover.turnRight(), nil
}
