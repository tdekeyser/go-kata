package marsrover

import (
	"reflect"
	"testing"
)

func TestForward(t *testing.T) {
	north, _ := Forward(NorthRover{0, 0})
	assertEqual(t, north, NorthRover{0, 1})

	south, _ := Forward(SouthRover{0, 0})
	assertEqual(t, south, SouthRover{0, 9})

	east, _ := Forward(EastRover{1, 1})
	assertEqual(t, east, EastRover{2, 1})

	west, _ := Forward(WestRover{0, 0})
	assertEqual(t, west, WestRover{9, 0})
}

func TestBackward(t *testing.T) {
	north, _ := Backward(NorthRover{0, 0})
	assertEqual(t, north, NorthRover{0, 9})

	south, _ := Backward(SouthRover{2, 3})
	assertEqual(t, south, SouthRover{2, 4})

	east, _ := Backward(EastRover{0, 0})
	assertEqual(t, east, EastRover{9, 0})

	west, _ := Backward(WestRover{0, 0})
	assertEqual(t, west, WestRover{1, 0})
}

func TestTurnLeft(t *testing.T) {
	rover, _ := Left(NorthRover{0, 0})
	_, northToWestOk := rover.(WestRover)
	assertEqual(t, northToWestOk, true)
}

func TestTurnRight(t *testing.T) {
	rover, _ := Right(NorthRover{0, 0})
	_, northToEastOk := rover.(EastRover)
	assertEqual(t, northToEastOk, true)
}

func TestGridControl(t *testing.T) {
	north, _ := Forward(NorthRover{0, 9})
	assertEqual(t, north, NorthRover{0, 0})

	east, _ := Backward(EastRover{0, 2})
	assertEqual(t, east, EastRover{9, 2})
}

func TestInstructions(t *testing.T) {
	north := NorthRover{0, 0}

	rover, _ := Perform(north, Instructions{
		Forward,
		Forward,
		Right,
		Forward,
	})

	assertEqual(t, rover, EastRover{1, 2})
}

func assertEqual(t *testing.T, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		t.Errorf("Not equal: %d %d", x, y)
	}
}
