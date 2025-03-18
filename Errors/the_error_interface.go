/*

The Error Interface
Go programs express errors with error values. An Error is any type that implements the simple built-in error interface:

type error interface {
    Error() string
}

When something can go wrong in a function, that function should return an error as its last return value. Any code that calls a function that can return an error should handle errors by testing whether the error is nil.

Atoi
Let's look at how the strconv.Atoi function uses this pattern. The function signature of Atoi is:

func Atoi(s string) (int, error)

This means Atoi takes a string argument and returns two values: an integer and an error. If the string can be successfully converted to an integer, Atoi returns the integer and a nil error. If the conversion fails, it returns zero and a non-nil error.

Here's how you would safely use Atoi:

// Atoi converts a stringified number to an integer
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
    // because "42b" isn't a valid integer, we print:
    // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
    // Note:
    // 'parsing "42b": invalid syntax' is returned by the .Error() method
    return
}
// if we get here, then the
// variable i was converted successfully

A nil error denotes success; a non-nil error denotes failure.

Assignment
We offer a product that allows businesses to send pairs of messages to couples. It is mostly used by flower shops and movie theaters.

Complete the sendSMSToCouple function. It should send 2 messages, first to the customer and then to the customer's spouse.

Use sendSMS() to send the msgToCustomer. If an error is encountered, return 0 and the error.
Do the same for the msgToSpouse
If both messages are sent successfully, return the total cost of the messages added together.
When you return a non-nil error in Go, it's conventional to return the "zero" values of all other return values.

*/

package main

import (
	"fmt"
	"testing"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	// ?
	toCustomerCost, error := sendSMS(msgToCustomer)
	if error != nil {
		return 0, error
	}

	toSpouseCost, error := sendSMS(msgToSpouse)
	if error != nil {
		return 0, error
	}

	return toCustomerCost + toSpouseCost, nil

}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func Test(t *testing.T) {
	type testCase struct {
		msgToCustomer string
		msgToSpouse   string
		expectedCost  int
		expectedErr   error
	}

	runCases := []testCase{
		{"Thanks for coming in to our flower shop today!", "We hope you enjoyed your gift.", 0, fmt.Errorf("can't send texts over 25 characters")},
		{"Thanks for joining us!", "Have a good day.", 76, nil},
	}

	submitCases := append(runCases, []testCase{
		{"Thank you.", "Enjoy!", 32, nil},
		{"We loved having you in!", "We hope the rest of your evening is fantastic.", 0, fmt.Errorf("can't send texts over 25 characters")},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		cost, err := sendSMSToCouple(test.msgToCustomer, test.msgToSpouse)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		expectedErrString := ""
		if test.expectedErr != nil {
			expectedErrString = test.expectedErr.Error()
		}
		if cost != test.expectedCost || errString != expectedErrString {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.msgToCustomer, test.msgToSpouse, test.expectedCost, test.expectedErr, cost, err)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.msgToCustomer, test.msgToSpouse, test.expectedCost, test.expectedErr, cost, err)
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
	Test(tests)
}
