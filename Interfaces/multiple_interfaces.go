/*

Multiple Interfaces
A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.

Assignment
Complete the required methods so that the email type implements both the expense and formatter interfaces.

cost()
If the email is not "subscribed", then the cost is 5 cents for each character in the body. If it is, then the cost is 2 cents per character.

Return the total cost of the entire email in cents.

format()
The format method should return a string in this format:

'CONTENT' | Subscribed

If the email is not subscribed, change the second part to "Not Subscribed":

'CONTENT' | Not Subscribed

The single quotes are included in the string, and CONTENT is the email's body. For example:

'Hello, World!' | Subscribed

Note: you may want to import the fmt package and use Sprintf.

*/

package main

import (
	"fmt"
)

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

func (e email) cost() int {
	// ?
	if e.isSubscribed == true {
		return 2 * len(e.body)
	}

	return 5 * len(e.body)
}

func (e email) format() string {
	if e.isSubscribed == true {
		return fmt.Sprintf("'%s' | Subscribed", e.body)
	}

	return fmt.Sprintf("'%s' | Not Subscribed", e.body)
}

func main() {
	type testCase struct {
		body           string
		isSubscribed   bool
		expectedCost   int
		expectedFormat string
	}

	runCases := []testCase{
		{"hello there", true, 22, "'hello there' | Subscribed"},
		{"general kenobi", false, 70, "'general kenobi' | Not Subscribed"},
	}

	submitCases := append(runCases, []testCase{
		{"i hate sand", true, 22, "'i hate sand' | Subscribed"},
		{"it's coarse and rough and irritating", false, 180, "'it's coarse and rough and irritating' | Not Subscribed"},
		{"and it gets everywhere", true, 44, "'and it gets everywhere' | Subscribed"},
	}...)

	testCases := runCases

	testCases = submitCases

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		e := email{
			body:         test.body,
			isSubscribed: test.isSubscribed,
		}
		cost := e.cost()
		format := e.format()
		if format != test.expectedFormat || cost != test.expectedCost {
			failCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.body, test.isSubscribed, test.expectedCost, test.expectedFormat, cost, format)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.body, test.isSubscribed, test.expectedCost, test.expectedFormat, cost, format)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}
