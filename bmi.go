package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikemando/bmi-calculator/name"
)

func main() {
	// Application info
	fmt.Println(name.AppName)
	fmt.Println(name.LineBreak)
	// Prompt user for Height and Weight input
	fmt.Print(name.HeightInput)
	userHeight, _ := inputReader.ReadString('\n')

	fmt.Print(name.WeightInput)
	userWeight, _ := inputReader.ReadString('\n')

	userHeight = strings.Replace(userHeight, "\n", "", -1)
	userWeight = strings.Replace(userWeight, "\n", "", -1)

	height, _ := strconv.ParseFloat(userHeight, 64)
	weight, _ := strconv.ParseFloat(userWeight, 64)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI is %.2f", bmi)
}
