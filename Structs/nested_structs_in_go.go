/*

Nested Structs in Go
Structs can be nested to represent more complex entities:

type car struct {
  brand string
  model string
  doors int
  mileage int
  frontWheel wheel
  backWheel wheel
}

type wheel struct {
  radius int
  material string
}

The fields of a struct can be accessed using the dot . operator.

myCar := car{}
myCar.frontWheel.radius = 5

Assignment
Textio has a bug, we've been sending texts that are missing critical bits of information! Before we send text messages in Textio, we must check to make sure the required fields have non-zero values.

Notice that the user struct is a nested struct within the messageToSend struct. Both sender and recipient are user struct types.

Complete the canSendMessage function. It should return true only if the sender and recipient fields each contain a name and a number. If any of the default zero values are present, return false instead.


*/

package main

import (
	"fmt"
	"testing"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 || mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func Test(t *testing.T) {
	type testCase struct {
		mToSend  messageToSend
		expected bool
	}

	runCases := []testCase{
		{messageToSend{
			message:   "you have an appointment tomorrow",
			sender:    user{name: "Brenda Halafax", number: 16545550987},
			recipient: user{name: "Sally Sue", number: 19035558973},
		}, true},
		{messageToSend{
			message:   "you have an event tomorrow",
			sender:    user{number: 16545550987},
			recipient: user{name: "Suzie Sall", number: 19035558973},
		}, false},
	}

	submitCases := append(runCases, []testCase{
		{messageToSend{
			message:   "you have an birthday tomorrow",
			sender:    user{name: "Jason Bjorn", number: 16545550987},
			recipient: user{name: "Jim Bond"},
		}, false},
		{messageToSend{
			message:   "you have a party tomorrow",
			sender:    user{name: "Njorn Halafax"},
			recipient: user{name: "Becky Sue", number: 19035558973},
		}, false},
		{messageToSend{
			message:   "you have a birthday tomorrow",
			sender:    user{name: "Eli Halafax", number: 16545550987},
			recipient: user{number: 19035558973},
		}, false},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := canSendMessage(test.mToSend)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:
  * message:          %s
  * sender.name:      %s
  * sender.number:    %d
  * recipient.name:   %s
  * recipient.number: %d
  Expected:           %v
  Actual:             %v
Fail
`,
				test.mToSend.message,
				test.mToSend.sender.name,
				test.mToSend.sender.number,
				test.mToSend.recipient.name,
				test.mToSend.recipient.number,
				test.expected,
				output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:
  * message:          %s
  * sender.name:      %s
  * sender.number:    %d
  * recipient.name:   %s
  * recipient.number: %d
  Expected:           %v
  Actual:             %v
Pass
`,
				test.mToSend.message,
				test.mToSend.sender.name,
				test.mToSend.sender.number,
				test.mToSend.recipient.name,
				test.mToSend.recipient.number,
				test.expected,
				output)
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
	mockTest := &testing.T{}

	// Run the test manually
	Test(mockTest)
}
