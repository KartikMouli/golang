/*

Variadic
Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the "..." syntax in the function signature.

A variadic function receives the variadic arguments as a slice.

func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}

The familiar fmt.Println() and fmt.Sprintf() are variadic! fmt.Println() prints each element with space delimiters and a newline at the end.

func Println(a ...interface{}) (n int, err error)

Spread Operator
The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}

Assignment
We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

Complete the sum function to return the sum of all inputs.

Take note of how the variadic inputs and the spread operator are used in the test suite.

*/

package main

import (
	"fmt"
)

func sum(nums ...int) int {
	// ?

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

func main() {
	type testCase struct {
		nums     []int
		expected int
	}

	runCases := []testCase{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 15},
	}

	submitCases := append(runCases, []testCase{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
		{[]int{}, 0},
		{[]int{5}, 5},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := sum(test.nums...)
		if output != test.expected {
			failCount++
			fmt.Printf(`---------------------------------
Inputs:
%v
Expecting:  %v
Actual:     %v
Fail
`, sliceWithBullets(test.nums), test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:
%v
Expecting:  %v
Actual:     %v
Pass
`, sliceWithBullets(test.nums), test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
