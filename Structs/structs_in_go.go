
/*

Structs in Go
We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:

type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
}

This creates a new struct type called car. All cars have a brand, model, doors and mileage.

Structs in Go are often used to represent data that you might use a dictionary or object for in other languages.

Assignment
Complete the definition of the messageToSend struct. It needs two fields:

phoneNumber - an integer
message - a string.


*/

package main

import (
	"fmt"
	"testing"
)


type messageToSend struct {
	phoneNumber int
	message string
}


func getMessageText(m messageToSend) string {
	return fmt.Sprintf("Sending message: '%s' to: %v", m.message, m.phoneNumber)
}

func Test(t *testing.T) {
	type testCase struct {
		phoneNumber int
		message     string
		expected    string
	}

	runCases := []testCase{
		{148255510981, "Thanks for signing up", "Sending message: 'Thanks for signing up' to: 148255510981"},
		{148255510982, "Love to have you aboard!", "Sending message: 'Love to have you aboard!' to: 148255510982"},
	}

	submitCases := append(runCases, []testCase{
		{148255510983, "We're so excited to have you", "Sending message: 'We're so excited to have you' to: 148255510983"},
		{148255510984, "", "Sending message: '' to: 148255510984"},
		{148255510985, "Hello, World!", "Sending message: 'Hello, World!' to: 148255510985"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMessageText(messageToSend{
			phoneNumber: test.phoneNumber,
			message:     test.message,
		})
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.phoneNumber, test.message, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.phoneNumber, test.message, test.expected, output)
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
	// Create a mock testing object
	mockTest := &testing.T{}

	// Run the test manually
	Test(mockTest)
}