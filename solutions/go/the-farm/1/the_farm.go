package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cowsCount int
	message   string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cowsCount, e.message)
}

func DivideFood(calculator FodderCalculator, cows int) (float64, error) {
	quantity, err := calculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	ff, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return float64(quantity) * ff / float64(cows), nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, cows int)(float64, error){
	if cows<1{
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calculator, cows)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 1 {
        message:="there are no negative cows"
       if cows == 0{
            message="no cows don't need food"
       }
		return &InvalidCowsError{
			cowsCount: cows,
			message:   message,
		}
	}
	return nil
}