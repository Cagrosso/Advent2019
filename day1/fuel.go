package main

import (
	"math"
)

// FuelCalculation determines amount of fuel needed for given mass
func FuelCalculation(mass int) int {
	var f64 float64 = float64(mass) / 3.0
	f64 = math.Floor(f64)

	return int(f64) - 2
}

// TotalFuelCalculation determines the amount of fuel needed to account for mass + mass of fuel
func TotalFuelCalculation(mass int) int {
	totalFuel := 0
	currentMass := mass
	for {
		currentMass = FuelCalculation(currentMass)

		if currentMass > 0 {
			totalFuel += currentMass
			continue
		} else {
			break
		}
	}

	return totalFuel
}

func main() {
	FuelCalculation(2)
}
