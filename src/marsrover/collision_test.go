package marsrover

import (
	"gotest.tools/assert"
	"testing"
)

func TestObstableDetection(t *testing.T) {
	east, err := Forward(EastRover{0, 0})

	assert.Error(t, err, "obstacle detected")
	assert.DeepEqual(t, east, EastRover{0, 0})
}

func TestInstructions_stopsWhenObstacleDetected(t *testing.T) {
	north := NorthRover{0, 0}

	rover, e := Perform(north, Instructions{
		Right,
		Forward,
		Forward,
	})

	assert.Error(t, e, "obstacle detected")
	assert.DeepEqual(t, rover, EastRover{0, 0})
}
