// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"formula/pkg/formula"
	"os"
	"strconv"
)

func main() {
	numberOne, _ := strconv.ParseFloat(os.Getenv("NUMBER_ONE"), 64)
	numberTwo, _ := strconv.ParseFloat(os.Getenv("NUMBER_TWO"), 64)

	formula.Formula{
		NumberOne: numberOne,
		NumberTwo: numberTwo,
	}.Run(os.Stdout)
}
