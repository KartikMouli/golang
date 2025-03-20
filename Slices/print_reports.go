/*

Print Reports
Anonymous functions are true to form in that they have no name.

Assignment
Complete the printReports function.

Call printCostReport once for each message. Pass in an anonymous function as the costCalculator that returns an int equal to twice the length of the input message.

*/

package main

import "fmt"

func printReports(messages []string) {
	// ?

	for _, msg := range messages {
		// Anonymous function to calculate cost
		costCalculator := func(message string) int {
			return 2 * len(message)
		}

		// Call printCostReport with the anonymous function and message
		printCostReport(costCalculator, msg)
	}
}

// don't touch below this line

func test(messages []string) {
	defer fmt.Println("====================================")
	printReports(messages)
}

func main() {
	test([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",
	})

	test([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",
	})
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
