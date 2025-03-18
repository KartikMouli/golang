/*

Connections
Textio has group chats that make communicating with multiple people much more efficient--if the chat doesn't descend into chaos. Instead of sending the message multiple times to individual people, you send one message to all of them at once.

Assignment
Complete the countConnections function that takes an integer groupSize representing the number of people in the group chat and returns an integer representing the number of connections between them. For each additional person in the group, the number of new connections continues to grow. Use a for loop to accumulate the number of connections instead of directly using a mathematical formula.

connections

*/

package main

import (
	"fmt"
)

func countConnections(groupSize int) int {
	// ?
	sum := 0

	for i := 0; i < groupSize; i++ {
		sum += i
	}

	return sum
}


func testCountGroupConnections() {
	type testCase struct {
		groupSize int
		expected  int
	}

	testCases := []testCase{
		{1, 0},
		{2, 1},
		{3, 3},
		{4, 6},
		{10, 45},
		{100, 4950},
		{1000, 499500},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result := countConnections(test.groupSize)
		if result != test.expected {
			failCount++
			fmt.Printf("FAIL - Group Size: %d | Expected: %d | Got: %d\n", test.groupSize, test.expected, result)
		} else {
			passCount++
			fmt.Printf("PASS - Group Size: %d | Expected: %d | Got: %d\n", test.groupSize, test.expected, result)
		}
	}

	fmt.Printf("\nResults: %d passed, %d failed\n", passCount, failCount)
}

func main() {
	testCountGroupConnections()
}