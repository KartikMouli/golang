/*

Formatting Strings Review
A convenient way to format strings in Go is by using the standard library's fmt.Sprintf() function. It's a string interpolation function, similar to Python's f-strings. The %v substring uses the type's default formatting, which is often what you want.

Default Values
const name = "Kim"
const age = 22
s := fmt.Sprintf("%v is %v years old.", name, age)
// s = "Kim is 22 years old."

The equivalent Python code:

name = "Kim"
age = 22
s = f"{name} is {age} years old."
# s = "Kim is 22 years old."

Rounding Floats
fmt.Printf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
fmt.Printf("I am %.2f years old", 10.523)
// I am 10.52 years old

Assignment
We need better error logs for our backend developers to help them debug their code.

Complete the getSMSErrorString() function. It should return a string with this format:

SMS that costs $COST to be sent to 'RECIPIENT' can not be sent

COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
RECIPIENT is the stringified representation of the recipient's phone number
Be sure to include the $ symbol and the single quotes

*/

package main

import (
	"fmt"
	"testing"
)

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%s' can not be sent",cost,recipient)
}




func Test(t *testing.T) {
	type testCase struct {
		cost      float64
		recipient string
		expected  string
	}

	runCases := []testCase{
		{1.4, "+1 (435) 555 0923", "SMS that costs $1.40 to be sent to '+1 (435) 555 0923' can not be sent"},
		{2.1, "+2 (702) 555 3452", "SMS that costs $2.10 to be sent to '+2 (702) 555 3452' can not be sent"},
	}

	submitCases := append(runCases, []testCase{
		{32.1, "+1 (801) 555 7456", "SMS that costs $32.10 to be sent to '+1 (801) 555 7456' can not be sent"},
		{14.4, "+1 (234) 555 6545", "SMS that costs $14.40 to be sent to '+1 (234) 555 6545' can not be sent"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getSMSErrorString(test.cost, test.recipient)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.cost, test.recipient, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.cost, test.recipient, test.expected, output)
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