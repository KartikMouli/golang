/*

Interface Implementation
Interfaces are implemented implicitly.

A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.

A quick way of checking whether a struct implements an interface is to declare a function that takes an interface as an argument. If the function can take the struct as an argument, then the struct implements the interface.

Assignment
At Textio we have full-time employees and contract employees. We have been tasked with making a more general employee interface so that dealing with different employee types is simpler.

Run the code. You should see an error indicating the contractor type does not fulfill the employee interface.

Add the missing getSalary method to the contractor type so that it fulfills the employee interface.

A contractor's salary is their hourly pay multiplied by how many hours they work per year.

*/

package main

import (
	"fmt"
	"testing"
)

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func Test(t *testing.T) {
	type testCase struct {
		emp      employee
		expected int
	}

	runCases := []testCase{
		{fullTime{name: "Bob", salary: 7300}, 7300},
		{contractor{name: "Jill", hourlyPay: 872, hoursPerYear: 982}, 856304},
	}

	submitCases := append(runCases, []testCase{
		{fullTime{name: "Alice", salary: 10000}, 10000},
		{contractor{name: "John", hourlyPay: 1000, hoursPerYear: 1000}, 1000000},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		salary := test.emp.getSalary()
		if salary != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  %v
Actual:     %v
Fail
`, test.emp, test.expected, salary)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %v
Actual:     %v
Pass
`, test.emp, test.expected, salary)
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
	tests := testing.T{}
	Test(&tests)
}
