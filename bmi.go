package main

import (
	"github.com/mikemando/bmi-calculator/name"
)

func main() {
	// Application info
	name.AppInfo()
	// Prompt user for Height and Weight input
	height, weight := userMetrics()

	bmi := calculateBMI(height, weight)

	displayBMI(bmi)
}

func calculateBMI(height float64, weight float64) float64 {
	return weight / (height * height)
}
