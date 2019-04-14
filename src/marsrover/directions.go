package marsrover

// In Go, operator % is the "remainder" after division, not the modulus.
// We need a function to get the modulus of x with divisor n.
func mod(x, n int) int {
	return ((x % n) + n) % n
}

// #####
// NORTH
// #####

type NorthRover struct {
	X, Y int
}

func (rover NorthRover) getPosition() (int, int) {
	return rover.X, rover.Y
}

func (rover NorthRover) move(amount int) Rover {
	return NorthRover{rover.X, mod(rover.Y+amount, Ymax)}
}

func (rover NorthRover) turnLeft() Rover {
	return WestRover{rover.X, rover.Y}
}

func (rover NorthRover) turnRight() Rover {
	return EastRover{rover.X, rover.Y}
}

// #####
// SOUTH
// #####

type SouthRover struct {
	X, Y int
}

func (rover SouthRover) getPosition() (int, int) {
	return rover.X, rover.Y
}

func (rover SouthRover) move(amount int) Rover {
	return SouthRover{rover.X, mod(rover.Y-amount, Ymax)}
}

func (rover SouthRover) turnLeft() Rover {
	return EastRover{rover.X, rover.Y}
}

func (rover SouthRover) turnRight() Rover {
	return WestRover{rover.X, rover.Y}
}

// ####
// EAST
// ####

type EastRover struct {
	X, Y int
}

func (rover EastRover) getPosition() (int, int) {
	return rover.X, rover.Y
}

func (rover EastRover) move(amount int) Rover {
	return EastRover{mod(rover.X+amount, Xmax), rover.Y}
}

func (rover EastRover) turnLeft() Rover {
	return NorthRover{rover.X, rover.Y}
}

func (rover EastRover) turnRight() Rover {
	return SouthRover{rover.X, rover.Y}
}

// ####
// WEST
// ####

type WestRover struct {
	X, Y int
}

func (rover WestRover) getPosition() (int, int) {
	return rover.X, rover.Y
}

func (rover WestRover) move(amount int) Rover {
	return WestRover{mod(rover.X-amount, Xmax), rover.Y}
}

func (rover WestRover) turnLeft() Rover {
	return SouthRover{rover.X, rover.Y}
}

func (rover WestRover) turnRight() Rover {
	return NorthRover{rover.X, rover.Y}
}
