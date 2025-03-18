/*

Type Assertions in Go
When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.

The example below shows how to safely access the radius field of s when s is an unknown type:

we want to check if s is a circle in order to cast it into its underlying concrete type
we know s is an instance of the shape interface, but we do not know if it's also a circle
c is a new circle struct cast from s
ok is true if s is indeed a circle, or false if s is NOT a circle
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

c, ok := s.(circle)
if !ok {
	// log an error if s isn't a circle
	log.Fatal("s is not a circle")
}

radius := c.radius

Assignment
Implement the getExpenseReport function.

If the expense is an email then it should return the email's toAddress and the cost of the email.
If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, just return an empty string and 0.0 for the cost.


*/

package main

import (
	"fmt"
)



var withSubmit = true


func getExpenseReport(e expense) (string, float64) {
	// ?
	// Check if the expense is an email
	if emailExpense, ok := e.(email); ok {
		return emailExpense.toAddress, emailExpense.cost()
	}

	// Check if the expense is an sms
	if smsExpense, ok := e.(sms); ok {
		return smsExpense.toPhoneNumber, smsExpense.cost()
	}

	// If it is neither, return empty string and 0.0
	return "", 0.0
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
			email{isSubscribed: true, body: "Whoa there!", toAddress: "soldier@monty.com"},
			"soldier@monty.com",
			0.11,
		},
		{
			sms{isSubscribed: false, body: "Halt! Who goes there?", toPhoneNumber: "+155555509832"},
			"+155555509832",
			2.1,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			email{
				isSubscribed: false,
				body:         "It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator of the Saxons, sovereign of all England!",
				toAddress:    "soldier@monty.com",
			},
			"soldier@monty.com",
			6.95,
		},
		{
			email{
				isSubscribed: true,
				body:         "Pull the other one!",
				toAddress:    "arthur@monty.com",
			},
			"arthur@monty.com",
			0.19,
		},
		{
			sms{
				isSubscribed:  true,
				body:          "I am. And this my trusty servant Patsy.",
				toPhoneNumber: "+155555509832",
			},
			"+155555509832",
			1.17,
		},
		{
			invalid{},
			"",
			0.0,
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	passCount := 0
	failCount := 0
	skipped := len(submitCases) - len(testCases)

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


