/*

Loops in Go
The basic loop in Go is written in standard C-like syntax:

for INITIAL; CONDITION; AFTER{
  // do something
}

INITIAL is run once at the beginning of the loop and can create
variables within the scope of the loop.

CONDITION is checked before each iteration. If the condition doesn't pass
then the loop breaks.

AFTER is run after each iteration.

For example:

for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9

Assignment
At Textio we have a dynamic formula for determining how much a batch of bulk messages costs to send. Complete the bulkSend() function.

It should return the total cost (as a float64) to send a batch of numMessages messages. Each message costs 1.0, plus an additional fee. The fee structure is:

1st message: 1.0 + 0.00
2nd message: 1.0 + 0.01
3rd message: 1.0 + 0.02
4th message: 1.0 + 0.03
...
Use a loop to calculate the total cost and return it.

*/



package main

import (
	"fmt"
	"testing"
)


func bulkSend(numMessages int) float64 {
	
	n := float64(numMessages)
	sum:=0.0

	for i:=0;i<numMessages;i++{
		sum+= float64(i)
	}

	sum/=100.0

	return n + sum
}



func Test(t *testing.T) {
	type testCase struct {
		numMessages int
		expected    float64
	}
	runCases := []testCase{
		{10, 10.45},
		{20, 21.9},
	}

	submitCases := append(runCases, []testCase{
		{0, 0.0},
		{1, 1.0},
		{5, 5.10},
		{30, 34.35},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := bulkSend(test.numMessages)
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Fail
`, test.numMessages, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Pass
`, test.numMessages, test.expected, output)
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
	Test(tests)
}
