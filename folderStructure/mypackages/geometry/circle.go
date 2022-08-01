package geometry

import "math"

func Ar(r float64) float64 {
	return math.Pi * r * r
}

func Per(r float64) float64 {
	return math.Pi * 2 * r
}