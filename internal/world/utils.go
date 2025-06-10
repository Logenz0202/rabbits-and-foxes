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
