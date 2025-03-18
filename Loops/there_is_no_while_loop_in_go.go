/*

There Is No While Loop in Go
Most programming languages have a concept of a while loop. Because Go allows for the omission of sections of a for loop, a while loop is just a for loop that only has a CONDITION.

for CONDITION {
  // do some stuff while CONDITION is true
}

For example:

plantHeight := 1
for plantHeight < 5 {
  fmt.Println("still growing! current height:", plantHeight)
  plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches")

Which prints:

still growing! current height: 1
still growing! current height: 2
still growing! current height: 3
still growing! current height: 4
plant has grown to 5 inches

Assignment
We have an interesting new cost structure from our SMS vendor. They charge exponentially more money for each consecutive text we send! Let's write a function that calculates how many messages we can send in a given batch given a costMultiplier and a maxCostInPennies.

In a nutshell, the first message costs a penny, and each message after that costs the same as the previous message multiplied by the costMultiplier.

There is a bug in the code! Add a condition to the for loop to fix the bug. The loop should stop when balance is equal to or less than 0. So what condition should the for loop evaluate?

*/

package main

import (
	"fmt"
	"testing"
)


func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance > 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}



func Test(t *testing.T) {
	type testCase struct {
		costMultiplier   float64
		maxCostInPennies int
		expected         int
	}

	runCases := []testCase{
		{1.1, 5, 4},
		{1.3, 10, 5},
		{1.35, 25, 7},
	}

	submitCases := append(runCases, []testCase{
		{1.2, 1, 1},
		{1.2, 15, 7},
		{1.3, 20, 7},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMaxMessagesToSend(test.costMultiplier, test.maxCostInPennies)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.costMultiplier, test.maxCostInPennies, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.costMultiplier, test.maxCostInPennies, test.expected, output)
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
	tests:= &testing.T{}
	Test(tests)
}