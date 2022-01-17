package robot_in_circle

// Robot starts at (0, 0) facing north
// G, straight 1
// L, 90 deg left
// R, 90 deg right
// Repeat forever
// return true if there exists a circle in the plane where the roboto never leaves

func RotateRobot(instruction rune, startingOrientation int) int {
	orientation := startingOrientation
	if instruction == 'L' {
		orientation += 90
	} else if instruction == 'R' {
		orientation -= 90
	} else {
		return startingOrientation
	}
	if orientation < 0 {
		orientation += 360
	}
	return orientation % 360
}

func MoveRobot(startingPos [2]int, orientation int, instruction rune) [2]int {
	if instruction != 'G' {
		return startingPos
	}

	// Move east
	if orientation == 90 {
		return [2]int{startingPos[0] + 1, startingPos[1]}
	// Move south
	} else if orientation == 180 {
		return [2]int{startingPos[0], startingPos[1] + 1}
		// Move west
	} else if orientation == 270 {
		return [2]int{startingPos[0] - 1, startingPos[1]}
	// Move north
	} else {
		return [2]int{startingPos[0], startingPos[1] - 1}
	}
}

func RobotInCircle(instructions string) bool {
	startingPos := [2]int{0, 0}
	orientation := 0 // Express orientiation in degrees
	pos := [2]int{0, 0}
	for _, instruction := range instructions {
		orientation = RotateRobot(instruction, orientation)
		pos = MoveRobot(pos, orientation, instruction)
	}
	return pos == startingPos || (orientation != 0 && orientation != 360)
}
