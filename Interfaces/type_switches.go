/*
Type Switches
A type switch makes it easy to do several type assertions in a series.

A type switch is similar to a regular switch statement, but the cases specify types instead of values.

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}

fmt.Printf("%T\n", v) prints the type of a variable.

Assignment
After submitting our last snippet of code for review, a more experienced gopher told us to use a type switch instead of successive assertions. Let's make that improvement!

Implement the getExpenseReport function using a type switch.

If the expense is an email then it should return the email's toAddress and the cost of the email.
If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, just return an empty string and 0.0 for the cost.

*/


package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	// ?
	switch v:= e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	type testCase struct {
		expense      expense
		expectedTo   string
		expectedCost float64
	}

	runCases := []testCase{
		{
			email{true, "hello there", "kit@boga.com"},
			"kit@boga.com",
			0.11,
		},
		{
			sms{false, "I'm a Nigerian prince, please send me your bank info so I can deposit $1000 dollars", "+155555509832"},
			"+155555509832",
			8.3,
		},
	}

	submitCases := append(runCases, []testCase{
		{invalid{}, "", 0},
		{
			email{false, "This meeting could have been an email", "jane@doe.com"},
			"jane@doe.com",
			1.85,
		},
		{
			sms{false, "Please sir/madam", "+155555504444"},
			"+155555504444",
			1.6,
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
		to, cost := getExpenseReport(test.expense)
		if to != test.expectedTo || cost != test.expectedCost {
			failCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
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
