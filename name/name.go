package name

import "fmt"

const appName = "BMI Calculator"
const lineBreak = "--------------------"
const HeightInput = "Please enter your height (m): "
const WeightInput = "Please enter your weight (kg): "

func AppInfo() {
	fmt.Println(appName)
	fmt.Println(lineBreak)
}
