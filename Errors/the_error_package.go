/*

The Errors Package
The Go standard library provides an "errors" package that makes it easy to deal with errors.

Read the godoc for the errors.New() function, but here's a simple example:

var err error = errors.New("something went wrong")

Assignment
Textio's software architects may have overcomplicated the requirements from the last coding assignment... oops. All we needed was a new generic error message that returns the string no dividing by 0 when a user attempts to get us to perform the taboo.

Complete the divide function. Use the errors.New() function to return an error when y == 0 that reads "no dividing by 0".

Hint
Remember that it's conventional to return the "zero" values of all other return values when you return a non-nil error in Go.

*/

package main

import (
	"errors"
	"fmt"
	"testing"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

func TestDivide(t *testing.T) {
	type testCase struct {
		x, y, expected float64
		expectedErr    string
	}

	runCases := []testCase{
		{10, 0, 0, "no dividing by 0"},
		{10, 2, 5, ""},
		{15, 30, 0.5, ""},
		{6, 3, 2, ""},
	}

	submitCases := append(runCases, []testCase{
		{0, 10, 0, ""},
		{100, 0, 0, "no dividing by 0"},
		{-10, -2, 5, ""},
		{-10, 2, -5, ""},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result, err := divide(test.x, test.y)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if result != test.expected || errString != test.expectedErr {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.x, test.y, test.expected, test.expectedErr, result, errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.x, test.y, test.expected, test.expectedErr, result, errString)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

func main() {
	tests := &testing.T{}
	TestDivide(tests)
}
