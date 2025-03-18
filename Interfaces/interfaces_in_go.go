/*

Interfaces in Go
Interfaces allow you to focus on what a type does rather than how it's built. They can help you write more flexible and reusable code by defining behaviors (like methods) that different types can share. This makes it easy to swap out or update parts of your code without changing everything else.

Interfaces are just collections of method signatures. A type "implements" an interface if it has methods that match the interface's method signatures.

In the following example, a "shape" must be able to return its area and perimeter. Both rect and circle fulfill the interface.

type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

When a type implements an interface, it can then be used as that interface type.

func printShapeData(s shape) {
	fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
}

Because we say the input is of type shape, we know that any argument must implement the .area() and .perimeter() methods.

As an example, because the empty interface doesn't require a type to have any methods at all, every type automatically implements the empty interface, written as:

interface{}

Assignment
The birthdayMessage and sendingReport structs already have implementations of the getMessage method. The getMessage method returns a string, and any type that implements the method can be considered a message (meaning it implements the message interface).

Add the getMessage() method signature as a requirement on the message interface.
Complete the sendMessage function. It should return:
The content of the message.
The cost of the message, which is the length of the message multiplied by 3.
Notice that your code doesn't care at all about whether a specific message is a birthdayMessage or a sendingReport!

*/

package main

import (
	"fmt"
	"testing"
	"time"
)

func sendMessage(msg message) (string, int) {
	messg := msg.getMessage()

	return messg,len(messg)*3
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func Test(t *testing.T) {
	type testCase struct {
		msg          message
		expectedText string
		expectedCost int
	}

	runCases := []testCase{
		{birthdayMessage{time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC), "John Doe"},
			"Hi John Doe, it is your birthday on 1994-03-21T00:00:00Z",
			168,
		},
		{sendingReport{"First Report", 10},
			`Your "First Report" report is ready. You've sent 10 messages.`,
			183,
		},
	}

	submitCases := append(runCases, []testCase{
		{birthdayMessage{time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC), "Bill Deer"},
			"Hi Bill Deer, it is your birthday on 1934-05-01T00:00:00Z",
			171,
		},
		{sendingReport{"Second Report", 20},
			`Your "Second Report" report is ready. You've sent 20 messages.`,
			186,
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
		text, cost := sendMessage(test.msg)
		if text != test.expectedText || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.msg, test.expectedText, test.expectedCost, text, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.msg, test.expectedText, test.expectedCost, text, cost)
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
	tests := testing.T{}
	Test(&tests)
}
