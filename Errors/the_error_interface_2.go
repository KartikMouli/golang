/*


The Error Interface
Because errors are just interfaces, you can build your own custom types that implement the error interface. Here's an example of a userError struct that implements the error interface:

type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}

It can then be used as an error:

func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}

Assignment
Our users frequently try to run custom analytics queries on their message deliverability metrics, and end up writing a bad query that tries to divide a number by zero. It's become such a problem that we need to make a new type of error for division by zero.

Update the code so that the divideError type implements the error interface. Its Error() method should just return a string formatted in the following way:

can not divide DIVIDEND by zero

Where DIVIDEND is the actual dividend of the divideError. Use the %v verb to format the dividend as a float.


*/


package main

import (
	"fmt"
	"testing"
)

type divideError struct {
	dividend float64
}

// ?

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func (d divideError) Error() string{
	return fmt.Sprintf("can not divide %v by zero",d.dividend)
}





func TestDivide(t *testing.T) {
	type testCase struct {
		dividend, divisor, expected float64
		expectedError               string
	}

	runCases := []testCase{
		{10, 2, 5, ""},
		{15, 3, 5, ""},
	}

	submitCases := append(runCases, []testCase{
		{10, 0, 0, "can not divide 10 by zero"},
		{15, 0, 0, "can not divide 15 by zero"},
		{100, 10, 10, ""},
		{16, 4, 4, ""},
		{30, 6, 5, ""},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output, err := divide(test.dividend, test.divisor)
		var errString string
		if err != nil {
			errString = err.Error()
		}
		if output != test.expected || errString != test.expectedError {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.dividend, test.divisor, test.expected, test.expectedError, output, errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.dividend, test.divisor, test.expected, test.expectedError, output, errString)
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

func main(){
	tests := &testing.T{}
	TestDivide(tests)
}