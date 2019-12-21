package main

import "math"

// FuelCalculation determins amount of fuel needed for given mass
func FuelCalculation(mass int) int {
	var f64 float64 = float64(mass) / 3.0
	f64 = math.Floor(f64)

	return int(f64) - 2
}

func main() {
	FuelCalculation(2)
}
