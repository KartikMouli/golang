/*

Range
Go provides syntactic sugar to iterate easily over elements of a slice:

for INDEX, ELEMENT := range SLICE {
}

The element is a copy of the value at that index.

For example:

fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape

Assignment
We need to be able to quickly detect bad words in the messages our system sends.

Complete the indexOfFirstBadWord function. If it finds any bad words in the message it should return the index of the first bad word in the msg slice. This will help us filter out naughty words from our messaging system. If no bad words are found, return -1 instead.

Use the range keyword.

Tip
Remember, you can ignore returned values with an underscore _ instead of creating an unused variable.

*/


package main

import (
	"fmt"
	
)

func indexOfFirstBadWord(msg []string, badWords []string) int {
	// ?

	for i,word := range msg{
		for _,s := range badWords{
			if word == s{
				return i
			}
		}
	} 
	return -1
}




func main() {
	type testCase struct {
		msg      []string
		badWords []string
		expected int
	}

	runCases := []testCase{
		{[]string{"hey", "there", "john"}, []string{"crap", "shoot", "frick", "dang"}, -1},
		{[]string{"ugh", "oh", "my", "frick"}, []string{"crap", "shoot", "frick", "dang"}, 3},
	}

	submitCases := append(runCases, []testCase{
		{[]string{"what", "the", "shoot", "I", "hate", "that", "crap"}, []string{"crap", "shoot", "frick", "dang"}, 2},
		{[]string{"crap", "shoot", "frick", "dang"}, []string{""}, -1},
		{[]string{""}, nil, -1},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := indexOfFirstBadWord(test.msg, test.badWords)
		if output != test.expected {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed:
message:
%v
bad words:
%v
Expecting:  %v
Actual:     %v
Fail
`, sliceWithBullets(test.msg), sliceWithBullets(test.badWords), test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
message:
%v
bad words:
%v
Expecting:  %v
Actual:     %v
Pass
`, sliceWithBullets(test.msg), sliceWithBullets(test.badWords), test.expected, output)
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