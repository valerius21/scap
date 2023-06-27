package main

import (
	"fmt"

	"github.com/valerius21/scap/pkg/fns"
)

func main() {
	points := 100_000_000
	estimate := fns.MathFn(points)

	fmt.Printf("Estimate of PI is %f\n", estimate)
}
