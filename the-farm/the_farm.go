package thefarm

import "errors"

func DivideFood(cowCount int, fodderCalculator FodderCalculator) (float64, error) {
	if cowCount == 0 {
		return 0, errors.New("something went wrong")
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

func ValidateInputAndDivideFood(cowCount int, fodderCalculator FodderCalculator) (float64, error) {
	if cowCount <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(cowCount, fodderCalculator)
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
