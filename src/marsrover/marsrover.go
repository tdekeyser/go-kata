package marsrover

type Rover interface {
	move(amount int) Rover
	turnLeft() Rover
	turnRight() Rover
}

func Forward(rover Rover) Rover {
	return rover.move(1)
}

func Backward(rover Rover) Rover {
	return rover.move(-1)
}

func Left(rover Rover) Rover {
	return rover.turnLeft()
}

func Right(rover Rover) Rover {
	return rover.turnRight()
}

// #####
// NORTH
// #####

type NorthRover struct {
	x, y int
}

func (rover NorthRover) move(amount int) Rover {
	return NorthRover{rover.x, rover.y + amount}
}

func (rover NorthRover) turnLeft() Rover {
	return WestRover{rover.x, rover.y}
}

func (rover NorthRover) turnRight() Rover {
	return EastRover{rover.x, rover.y}
}

// #####
// SOUTH
// #####

type SouthRover struct {
	x, y int
}

func (rover SouthRover) move(amount int) Rover {
	return SouthRover{rover.x, rover.y - amount}
}

func (rover SouthRover) turnLeft() Rover {
	return EastRover{rover.x, rover.y}
}

func (rover SouthRover) turnRight() Rover {
	return WestRover{rover.x, rover.y}
}

// ####
// EAST
// ####

type EastRover struct {
	x, y int
}

func (rover EastRover) move(amount int) Rover {
	return EastRover{rover.x + amount, rover.y}
}

func (rover EastRover) turnLeft() Rover {
	return NorthRover{rover.x, rover.y}
}

func (rover EastRover) turnRight() Rover {
	return SouthRover{rover.x, rover.y}
}

// ####
// WEST
// ####

type WestRover struct {
	x, y int
}

func (rover WestRover) move(amount int) Rover {
	return WestRover{rover.x - amount, rover.y}
}

func (rover WestRover) turnLeft() Rover {
	return SouthRover{rover.x, rover.y}
}

func (rover WestRover) turnRight() Rover {
	return NorthRover{rover.x, rover.y}
}
