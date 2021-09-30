package darts

import "math"

// Calculate Pythagorean theorem
func Pythagoras(x float64, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2.0) + math.Pow(y, 2.0))
}

func Score(x float64, y float64) int {
	distance := Pythagoras(x, y)

	switch {
	case distance <= 1.0:
		return 10
	case distance <= 5.0:
		return 5
	case distance <= 10.0:
		return 1
	default:
		return 0

	}
}
