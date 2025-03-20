/*

Make
Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:

// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)

Slices created with make will be filled with the zero value of the type.

If we want to create a slice with a specific set of values, we can use a slice literal:

mySlice := []string{"I", "love", "go"}

Notice the square brackets do not have a 3 in them. If they did, you'd have an array instead of a slice.

Length
The length of a slice is simply the number of elements it contains. It is accessed using the built-in len() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(len(mySlice)) // 3

Capacity
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. It is accessed using the built-in cap() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3

Generally speaking, unless you're hyper-optimizing the memory usage of your program, you don't need to worry about the capacity of a slice because it will automatically grow as needed.

Assignment
We send a lot of text messages at Textio, and our API is getting slow and unresponsive.

If we know the rough size of a slice before we fill it up, we can make our program faster by creating the slice with that size ahead of time so that the Go runtime doesn't need to continuously allocate new underlying arrays of larger and larger sizes. By setting the length, the slice can still be resized later, but it means we can avoid all the expensive resizing since we know what we'll need.

Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.

Preallocate a slice for the message costs of the same length as the messages slice.
Fill the costs slice with costs for each message. The cost in the cost slice should correspond to the message in the messages slice at the same index. The cost of a message is the length of the message multiplied by 0.01.

*/

package main

import (
	"fmt"
	
)

func getMessageCosts(messages []string) []float64 {
	// ?
	cost := make([]float64, len(messages))

	for i := 0; i < len(messages); i++ {
		cost[i] = float64(len(messages[i])) * 0.01
	}

	return cost

}

func main() {
	type testCase struct {
		messages    []string
		expected    []float64
		expectedCap int
	}

	runCases := []testCase{
		{
			[]string{"Welcome to the movies!", "Enjoy your popcorn!"},
			[]float64{0.22, 0.19},
			2,
		},
		{
			[]string{"I don't want to be here anymore", "Can we go home?", "I'm hungry", "I'm bored"},
			[]float64{0.31, 0.15, 0.1, 0.09},
			4,
		},
	}

	submitCases := append(runCases, []testCase{
		{[]string{}, []float64{}, 0},
		{[]string{""}, []float64{0}, 1},
		{[]string{"Hello", "Hi", "Hey"}, []float64{0.05, 0.02, 0.03}, 3},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMessageCosts(test.messages)
		if !slicesEqual(output, test.expected) || cap(output) != test.expectedCap {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed:
%v
Expecting:
%v
expected cap: %v
Actual:
%v
actual cap: %v
Fail
`, sliceWithBullets(test.messages), sliceWithBullets(test.expected), test.expectedCap, sliceWithBullets(output), cap(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
%v
Expecting:
%v
expected cap: %v
Actual:
%v
actual cap: %v
Pass
`, sliceWithBullets(test.messages), sliceWithBullets(test.expected), test.expectedCap, sliceWithBullets(output), cap(output))
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

func slicesEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
