/*

Arrays in Go
Arrays are fixed-size groups of variables of the same type. For example, [4]string is an array of 4 values of type string.

To declare an array of 10 integers:

var myInts [10]int

or to declare an initialized literal:

primes := [6]int{2, 3, 5, 7, 11, 13}

Assignment
When our clients don't respond to a message, they can be reminded with up to 2 additional messages.

Complete the getMessageWithRetries function. It takes three strings and returns:

An array of 3 strings
An array of 3 integers
The returned string array contains the original messages. The first is the primary message, the second is the first reminder, and the third is the last reminder.

The integers in the integer array represent the cost of sending each message. The cost of each message is equal to the length of the message, plus the length of any previous messages. For example:

"hello" costs 5
"world" costs 5, adding "hello" makes total cost 10 (5 + 5)
"!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)

*/

package main

import (
	"fmt"
	
)

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	// ?
	var msg [3]string
	var cost [3]int

	msg[0] = primary
	msg[1] = secondary
	msg[2] = tertiary

	cost[0] = len(primary)
	cost[1] = cost[0] + len(secondary)
	cost[2] = cost[1] + len(tertiary)

	return msg, cost
}

func main() {
	type testCase struct {
		messages         []string
		expectedMessages [3]string
		expectedCosts    [3]int
	}

	runCases := []testCase{
		{
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]int{46, 92, 119},
		},
		{
			[]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]int{27, 49, 61},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			[]string{
				"Put that coffee down!",
				"Coffee is for closers",
				"Always be closing",
			},
			[3]string{
				"Put that coffee down!",
				"Coffee is for closers",
				"Always be closing",
			},
			[3]int{21, 42, 59},
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actualMessages, actualCosts := getMessageWithRetries(test.messages[0], test.messages[1], test.messages[2])
		if actualMessages[0] != test.expectedMessages[0] ||
			actualMessages[1] != test.expectedMessages[1] ||
			actualMessages[2] != test.expectedMessages[2] ||
			actualCosts[0] != test.expectedCosts[0] ||
			actualCosts[1] != test.expectedCosts[1] ||
			actualCosts[2] != test.expectedCosts[2] {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed:
Inputs:
%v
Expecting:
%v
%v
Actual:
%v
%v
Fail
`, sliceWithBullets(test.messages), sliceWithBullets(test.expectedMessages[:]), test.expectedCosts, sliceWithBullets(actualMessages[:]), actualCosts)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Inputs:
%v
Expecting:
%v
%v
Actual:
%v
%v
Pass
`, sliceWithBullets(test.messages), sliceWithBullets(test.expectedMessages[:]), test.expectedCosts, sliceWithBullets(actualMessages[:]), actualCosts)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

func sliceWithBullets[T any](slice []T) string {
	output := ""
	for i, item := range slice {
		form := "  - %v\n"
		if i == (len(slice) - 1) {
			form = "  - %v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true


