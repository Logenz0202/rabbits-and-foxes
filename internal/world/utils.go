package world

// Clamp restricts a given integer value v to lie within the range defined by min and max inclusively.

func Clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Abs returns the absolute value of the given integer x.

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IsAdjacent checks if the points (x1, y1) and (x2, y2) are adjacent on a grid,
// including diagonally, without being the same point.

func IsAdjacent(x1, y1, x2, y2 int) bool {
	dx := Abs(x1 - x2)
	dy := Abs(y1 - y2)
	return dx <= 1 && dy <= 1 && !(dx == 0 && dy == 0)
}
