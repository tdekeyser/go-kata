package marsrover

import (
	"gotest.tools/assert"
	"testing"
)

func TestForward(t *testing.T) {
	assert.DeepEqual(t, Forward(NorthRover{0, 0}), NorthRover{0, 1})
	assert.DeepEqual(t, Forward(SouthRover{0, 0}), SouthRover{0, 9})
	assert.DeepEqual(t, Forward(EastRover{1, 1}), EastRover{2, 1})
	assert.DeepEqual(t, Forward(WestRover{0, 0}), WestRover{9, 0})
}

func TestBackward(t *testing.T) {
	assert.DeepEqual(t, Backward(NorthRover{0, 0}), NorthRover{0, 9})
	assert.DeepEqual(t, Backward(SouthRover{2, 3}), SouthRover{2, 4})
	assert.DeepEqual(t, Backward(EastRover{0, 0}), EastRover{9, 0})
	assert.DeepEqual(t, Backward(WestRover{0, 0}), WestRover{1, 0})
}

func TestTurnLeft(t *testing.T) {
	_, northToWestOk := Left(NorthRover{0, 0}).(WestRover)
	assert.Assert(t, northToWestOk)
}

func TestTurnRight(t *testing.T) {
	_, northToEastOk := Right(NorthRover{0, 0}).(EastRover)
	assert.Assert(t, northToEastOk)
}

func TestCombination(t *testing.T) {
	assert.DeepEqual(t,
		Forward(Right(Backward(Right(Forward(NorthRover{0, 0}))))),
		SouthRover{9, 0})
}

func TestGridControl(t *testing.T) {
	assert.Equal(t, Forward(NorthRover{1, 9}), NorthRover{1, 0})
	assert.Equal(t, Backward(EastRover{0, 2}), EastRover{9, 2})
}
