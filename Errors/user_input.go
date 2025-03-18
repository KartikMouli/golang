/*
User Input
In Textio, users can set their profile status to communicate their current activity to those that choose to read it... However, there are some restrictions on what these statuses can contain. Your task is to implement a function that validates a user's status update. The status update cannot be empty and must not exceed 140 characters.

Assignment
Complete the validateStatus function. It should return an error when the status update violates any of the rules:

If the status is empty, return an error that reads status cannot be empty.
If the status exceeds 140 characters, return an error that says status exceeds 140 characters.
*/
package main

import (
	"errors"
	"fmt"
	"testing"
)

func validateStatus(status string) error {
	// ?
	if status == "" {
		return errors.New("status cannot be empty")
	}

	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}

	return nil

}

func TestValidateStatus(t *testing.T) {
	type testCase struct {
		status      string
		expectedErr string
	}

	runCases := []testCase{
		{"", "status cannot be empty"},
		{"This is a valid status update that is well within the character limit.", ""},
		{"This status update is way too long. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.", "status exceeds 140 characters"},
	}

	submitCases := append(runCases, []testCase{
		{"Another valid status.", ""},
		{"This status update, while derivative, contains exactly one hundred and forty-one characters, which is over the status update character limit.", "status exceeds 140 characters"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		err := validateStatus(test.status)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if errString != test.expectedErr {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     "%v"
Expecting:  "%v"
Actual:     "%v"
Fail
`, test.status, test.expectedErr, errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     "%v"
Expecting:  "%v"
Actual:     "%v"
Pass
`, test.status, test.expectedErr, errString)
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
	TestValidateStatus(tests)
}
