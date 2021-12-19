package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mikemando/bmi-calculator/name"
)

var inputReader = bufio.NewReader(os.Stdin)

func userMetrics() (float64, float64) {
	height := getUserInput(name.HeightInput)
	weight := getUserInput(name.WeightInput)

	return height, weight
}

func getUserInput(enteredText string) (value float64) {
	fmt.Print(enteredText)

	enteredValue, _ := inputReader.ReadString('\n')
	enteredValue = strings.Replace(enteredValue, "\n", "", -1)
	value, _ = strconv.ParseFloat(enteredValue, 64)

	return
}
