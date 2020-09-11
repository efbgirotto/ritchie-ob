// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io"

	"github.com/gookit/color"
)

type Formula struct {
	NumberOne float64
	NumberTwo float64
}

func (f Formula) Run(writer io.Writer) {
	var result string
	sum := f.NumberOne + f.NumberTwo

	result += fmt.Sprintf("Level 1: Basic Inputs!\n\n")

	result += color.FgGreen.Render(fmt.Sprintf("%v + %v = %v.\n", f.NumberOne, f.NumberTwo, sum))

	if _, err := fmt.Fprintf(writer, result); err != nil {
		panic(err)
	}
}
