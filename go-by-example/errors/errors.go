package main

import (
	"errors"
	"fmt"
)

type argumentError struct {
	arg     int
	message string
}

func (e *argumentError) Error() string { //if argument error implements the error interface it is of type error
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// no exemptions, return error as the last return value as convention
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argumentError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

// can assign errors to variables
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(numberOfCups int) error {
	maxCups := 3
	if numberOfCups > maxCups {
		return ErrOutOfTea
	} else if numberOfCups == 0 {
		return fmt.Errorf("this is an error wrapped in another error: %w", ErrPower) //we can wrap errors with higher level errors to give more context
	}

	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil { //can use inline error checks like this
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	_, err := f(42)

	argumentError, ok := errors.AsType[*argumentError](err) //as type checks the type and converts like is and as in c#
	if ok {
		fmt.Println(argumentError.arg)
		fmt.Println(argumentError.message)
	}

	err = makeTea(10)

	if errors.Is(err, ErrOutOfTea) { // can check the type of error with errors.Is()
		fmt.Println("We are out of tea!!!")
	}
}
