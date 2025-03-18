/*
Struct Methods in Go


While Go is not object-oriented, it does support methods that can be defined on structs. Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function.

type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
  return r.width * r.height
}

var r = rect{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
// prints 50

A receiver is just a special kind of function parameter. In the example above, the r in (r rect) could just as easily have been rec or even x, y or z. By convention, Go code will often use the first letter of the struct's name.

Receivers are important because they will, as you'll learn in the exercises to come, allow us to define interfaces that our structs (and other types) can implement.

Assignment
Let's clean up Textio's authentication logic. We store our user's authentication data inside an authenticationInfo struct. We need a method that can take that data and return a basic authorization string.

The format of the string should be:

Authorization: Basic USERNAME:PASSWORD

Create a method on the authenticationInfo struct called getBasicAuth that returns the formatted string.

*/

package main

import (
	"fmt"
	"testing"
)

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func TestGetBasicAuth(t *testing.T) {
	type testCase struct {
		auth     authenticationInfo
		expected string
	}

	runCases := []testCase{
		{authenticationInfo{"Google", "12345"}, "Authorization: Basic Google:12345"},
		{authenticationInfo{"Bing", "98765"}, "Authorization: Basic Bing:98765"},
	}

	submitCases := append(runCases, []testCase{
		{authenticationInfo{"DDG", "76921"}, "Authorization: Basic DDG:76921"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := test.auth.getBasicAuth()
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Fail
`, test.auth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Pass
`, test.auth, test.expected, output)
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
	TestGetBasicAuth(mockTest)
}
