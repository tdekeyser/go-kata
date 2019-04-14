package marsrover

import (
	"gotest.tools/assert"
	"testing"
)

func TestForward(t *testing.T) {
	north, _ := Forward(NorthRover{0, 0})
	assert.DeepEqual(t, north, NorthRover{0, 1})

	south, _ := Forward(SouthRover{0, 0})
	assert.DeepEqual(t, south, SouthRover{0, 9})

	east, _ := Forward(EastRover{1, 1})
	assert.DeepEqual(t, east, EastRover{2, 1})

	west, _ := Forward(WestRover{0, 0})
	assert.DeepEqual(t, west, WestRover{9, 0})
}

func TestBackward(t *testing.T) {
	north, _ := Backward(NorthRover{0, 0})
	assert.DeepEqual(t, north, NorthRover{0, 9})

	south, _ := Backward(SouthRover{2, 3})
	assert.DeepEqual(t, south, SouthRover{2, 4})

	east, _ := Backward(EastRover{0, 0})
	assert.DeepEqual(t, east, EastRover{9, 0})

	west, _ := Backward(WestRover{0, 0})
	assert.DeepEqual(t, west, WestRover{1, 0})
}

func TestTurnLeft(t *testing.T) {
	rover, _ := Left(NorthRover{0, 0})
	_, northToWestOk := rover.(WestRover)
	assert.Assert(t, northToWestOk)
}

func TestTurnRight(t *testing.T) {
	rover, _ := Right(NorthRover{0, 0})
	_, northToEastOk := rover.(EastRover)
	assert.Assert(t, northToEastOk)
}

func TestGridControl(t *testing.T) {
	north, _ := Forward(NorthRover{0, 9})
	assert.DeepEqual(t, north, NorthRover{0, 0})

	east, _ := Backward(EastRover{0, 2})
	assert.DeepEqual(t, east, EastRover{9, 2})
}

func TestInstructions(t *testing.T) {
	north := NorthRover{0, 0}

	rover, e := Perform(north, Instructions{
		Forward,
		Forward,
		Right,
		Forward,
	})

	assert.NilError(t, e)
	assert.DeepEqual(t, rover, EastRover{1, 2})
}
