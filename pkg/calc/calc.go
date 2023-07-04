package calc

import "fmt"

type calculator struct {
	// a, b float64 // операнды
	// op   string  // оператор
}

func New() calculator {
	return calculator{}
}

func (c *calculator) sum(a, b float64) float64 {
	return a + b
}

func (c *calculator) subt(a, b float64) float64 {
	return a - b
}

func (c *calculator) mul(a, b float64) float64 {
	return a * b
}

func (c *calculator) div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Calculate возвращает результат операции op с операндами a, b
func (c *calculator) Calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return c.sum(a, b), nil
	case "-":
		return c.subt(a, b), nil
	case "*":
		return c.mul(a, b), nil
	case "/":
		return c.div(a, b)
	default:
		return 0, fmt.Errorf("wrong operator")
	}
}
