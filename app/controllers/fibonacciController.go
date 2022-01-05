package controllers

import (
	"fmt"
	"math/big"

	"github.com/revel/revel"
)

type FibonacciController struct {
	*revel.Controller
}

func (c FibonacciController) Fibonacci() revel.Result {
	result := "[0, 1"
	count := 5000
	n1 := big.NewInt(0)
	n2 := big.NewInt(1)

	for i := 2; i < count; i++ {
		n1.Add(n1, n2)
		result = result + ", " + fmt.Sprint(n1)
		n1, n2 = n2, n1
	}

	result = result + "]"

	return c.RenderText(result)
}
