package marsrover

import (
	"fmt"
	"testing"
)

func TestObstableDetection(t *testing.T) {
	east, err := Forward(EastRover{0, 0})

	assertError(t, err, "obstacle detected")
	assertEqual(t, east, EastRover{0, 0})
}

func TestInstructions_stopsWhenObstacleDetected(t *testing.T) {
	north := NorthRover{0, 0}

	rover, e := Perform(north, Instructions{
		Right,
		Forward,
		Forward,
	})

	assertError(t, e, "obstacle detected")
	assertEqual(t, rover, EastRover{0, 0})
}

func assertError(t *testing.T, e error, msg string) {
	switch {
	case e == nil:
		t.Errorf("expected an error, got nil")
	case e.Error() != msg:
		t.Errorf(fmt.Sprintf("expected error %q, got %s", msg, e))
	}
}
