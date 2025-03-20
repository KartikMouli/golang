/*

Slices in Go
99 times out of 100 you will use a slice instead of an array when working with ordered lists.

Arrays are fixed in size. Once you make an array like [10]int you can't add an 11th element.

A slice is a dynamically-sized, flexible view of the elements of an array.

The zero value of slice is nil.

Non-nil slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:

primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}

The syntax is:

arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]

Where lowIndex is inclusive and highIndex is exclusive.

lowIndex, highIndex, or both can be omitted to use the entire array on that side of the colon.

Assignment
Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input as well as an array of 3 messages. You've been provided with constants representing the plan types at the top of the file.

If the plan is a pro plan, return all the strings from the messages input in a slice.
If the plan is a free plan, return the first 2 strings from the messages input in a slice.
If the plan isn't either of those, return a nil slice and an error that says unsupported plan.


*/

package main

import (
	"errors"
	"fmt"
	"slices"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	switch plan {
	case planFree:
		return messages[:2], nil
	case planPro:
		return messages[:], nil
	default:
		return nil, errors.New("unsupported plan")
	}
}

func main() {
	type testCase struct {
		plan             string
		messages         [3]string
		expectedMessages []string
		expectedErr      string
	}
	runCases := []testCase{
		{
			planFree,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{"Hello sir/madam can I interest you in a yacht?", "Please I'll even give you an Amazon gift card?"},
			"",
		},
		{
			planPro,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			"",
		},
	}

	submitCases := append(runCases, []testCase{
		{
			planFree,
			[3]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			[]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
			},
			"",
		},
		{
			planPro,
			[3]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			[]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			"",
		},
		{
			"invalid plan",
			[3]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			nil,
			"unsupported plan",
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
		actualMessages, err := getMessageWithRetriesForPlan(test.plan, test.messages)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if !slices.Equal(actualMessages, test.expectedMessages) || errString != test.expectedErr {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed:
Plan: %v
Messages:
%v
Expecting:
%v
errString:  %v
Actual:
%v
errString:  %v
Fail
`, test.plan, sliceWithBullets(test.messages[:]), sliceWithBullets(test.expectedMessages), test.expectedErr, sliceWithBullets(actualMessages), errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Plan: %v
Messages:
%v
Expecting:
%v
errString:  %v
Actual:
%v
errString:  %v
Pass
`, test.plan, sliceWithBullets(test.messages[:]), sliceWithBullets(test.expectedMessages), test.expectedErr, sliceWithBullets(actualMessages), errString)
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
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
