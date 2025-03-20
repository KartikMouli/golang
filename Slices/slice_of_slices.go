/*

Slice of Slices
Slices can hold other slices, effectively creating a matrix, or a 2D slice.

rows := [][]int{}

Assignment
We support various visualization dashboards on Textio that display message analytics to our users. The UI for our graphs and charts is built on top of a grid system. Let's build some grid logic.

Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers. The value of each cell is i * j where i and j are the indexes of the row and column respectively. Basically, we're building a multiplication chart.

For example, a 5x10 matrix, produced from calling createMatrix(5, 10), would look like this:

[0  0  0  0  0  0  0  0  0  0]
[0  1  2  3  4  5  6  7  8  9]
[0  2  4  6  8 10 12 14 16 18]
[0  3  6  9 12 15 18 21 24 27]
[0  4  8 12 16 20 24 28 32 36]


*/

package main

import (
	"fmt"
	"reflect"
)

func createMatrix(rows, cols int) [][]int {
	// ?
	mat := make([][]int, rows)

	for i := 0; i < rows; i++ {
		
		for j := 0; j < cols; j++ {
			mat[i] = append(mat[i], i*j)
		}
	}

	return mat
}

func main() {
	type testCase struct {
		rows, cols int
		expected   [][]int
	}

	runCases := []testCase{
		{3, 3, [][]int{
			{0, 0, 0},
			{0, 1, 2},
			{0, 2, 4},
		}},
		{4, 4, [][]int{
			{0, 0, 0, 0},
			{0, 1, 2, 3},
			{0, 2, 4, 6},
			{0, 3, 6, 9},
		}},
	}

	submitCases := append(runCases, []testCase{
		{5, 7, [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 2, 3, 4, 5, 6},
			{0, 2, 4, 6, 8, 10, 12},
			{0, 3, 6, 9, 12, 15, 18},
			{0, 4, 8, 12, 16, 20, 24},
		}},
		{0, 0, [][]int{}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := createMatrix(test.rows, test.cols)
		if !reflect.DeepEqual(output, test.expected) {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed: %v x %v matrix
Expecting:
%v
Actual:
%v
Fail
`, test.rows, test.cols, formatMatrix(test.expected), formatMatrix(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed: %v x %v matrix
Expecting:
%v
Actual:
%v
Pass
`, test.rows, test.cols, formatMatrix(test.expected), formatMatrix(output))
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

func formatMatrix(matrix [][]int) string {
	var result string
	for _, row := range matrix {
		result += fmt.Sprintf("%v\n", row)
	}
	return result
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
