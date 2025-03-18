/*

Omitting Conditions from a for Loop in Go
Loops in Go can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.

for INITIAL; ; AFTER {
  // do something forever
}

Assignment
Complete the maxMessages function. Given a cost threshold, it should calculate the maximum number of messages that can be sent.

Each message costs 100 pennies, plus an additional fee. The fee structure is:

1st message: 100 + 0
2nd message: 100 + 1
3rd message: 100 + 2
4th message: 100 + 3
Browser Freeze
If you lock up your browser by creating an infinite loop that isn't breaking, just click the cancel button.

*/

package main

import (
	"fmt"
)

func maxMessages(thresh int) int {
	// ?
	sum := 0
	for i := 0; ; i++ {
		if sum > thresh {
			return i - 1
		}
		sum += 100 + i
	}

}

func main() {
	testCases := []struct {
		thresh   int
		expected int
	}{
		{103, 1},
		{205, 2},
		{1000, 9},
		{100, 1},
		{3000, 26},
		{4000, 34},
		{5000, 41},
		{0, 0},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := maxMessages(test.thresh)
		if output != test.expected {
			failCount++
			fmt.Printf("FAIL: Inputs: (%v) | Expected: %v | Got: %v\n", test.thresh, test.expected, output)
		} else {
			passCount++
			fmt.Printf("PASS: Inputs: (%v) | Expected: %v | Got: %v\n", test.thresh, test.expected, output)
		}
	}

	fmt.Printf("\n%d passed, %d failed\n", passCount, failCount)
}
