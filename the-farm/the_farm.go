package thefarm

import "errors"

// TODO: define the 'DivideFood' function
func DivideFood(cowCount int, fodderCalculator FodderCalculator) (float64, error) {
	if cowCount <= 0 {
		return 0, errors.New("don't divide food by zero cow")
	}
	fodderAmount, err := fodderCalculator.FodderAmount(cowCount)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fodderAmount * fatteningFactor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function

// TODO: define the 'ValidateNumberOfCows' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
