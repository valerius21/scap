package fns

import (
	"math"
	"math/rand"
	"time"
)

// estimatePI estimates the value of PI using the Monte Carlo method.
func estimatePI(numPoints int) float64 {
	rand.Seed(time.Now().UnixNano())

	pointsInsideCircle := 0

	for i := 0; i < numPoints; i++ {
		x := rand.Float64()
		y := rand.Float64()
		distance := math.Sqrt(x*x + y*y)
		if distance <= 1 {
			pointsInsideCircle++
		}
	}

	pi := 4.0 * float64(pointsInsideCircle) / float64(numPoints)
	return pi
}

// MathFn is a function that estimates the value of PI using the Monte Carlo method.
func MathFn(points int) float64 {
	return estimatePI(points)
}
