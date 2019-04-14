package marsrover

var obstacles = [][]int{
	{1, 0}, {3, 3}, {5, 9},
}

func collidesWithObstacle(rover Rover) bool {
	for _, obstacle := range obstacles {
		x, y := rover.getPosition()
		if obstacle[0] == x && obstacle[1] == y {
			return true
		}
	}
	return false
}
