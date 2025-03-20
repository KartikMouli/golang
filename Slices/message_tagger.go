/*

Message Tagger
Textio needs a way to tag messages based on specific criteria.

Assignment
Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes a sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.
It should loop through each message and set the tags to the result of the passed in function.
Be sure to modify the messages of the original slice.
Ensure the tags field contains an initialized slice. No nil slices.
Complete the tagger function. It should take an sms message and return a slice of strings.
For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.
Regardless of casing just means that even "uRGent" or "SALE" should trigger the tag.

Example usage:

messages := []sms{
	{id: "001", content: "Urgent! Last chance to see!"},
	{id: "002", content: "Big sale on all items!"},
	// Additional messages...
}
taggedMessages := tagMessages(messages, tagger)
// `taggedMessages` will now have tags based on the content.
// 001 = [Urgent]
// 002 = [Promo]

Tip
The go strings package, specifically the Contains and ToLower functions, can be very helpful here!

*/

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// ?
	for i := range messages {
		messages[i].tags = tagger(messages[i])

	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	// ?
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}

	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

func main() {
	type testCase struct {
		messages []sms
		expected [][]string
	}

	runCases := []testCase{
		{
			messages: []sms{{id: "001", content: "Urgent, please respond!"}, {id: "002", content: "Big sale on all items!"}},
			expected: [][]string{{"Urgent"}, {"Promo"}},
		},
		{
			messages: []sms{{id: "003", content: "Enjoy your day"}},
			expected: [][]string{{}},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			messages: []sms{{id: "004", content: "Sale! Don't miss out on these urgent promotions!"}},
			expected: [][]string{{"Urgent", "Promo"}},
		},
		{
			messages: []sms{{id: "005", content: "i nEEd URgEnt help, my FROZEN FLAME was used"}, {id: "006", content: "wAnt to saLE 200x heavy leather"}},
			expected: [][]string{{"Urgent"}, {"Promo"}},
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
		actual := tagMessages(test.messages, tagger)
		if len(actual) != len(test.expected) {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed for length of returned sms slice
Expecting: %v
Actual:    %v
Fail
`, len(test.expected), len(actual))
			continue
		}

		for i, msg := range actual {
			if !reflect.DeepEqual(msg.tags, test.expected[i]) {
				failCount++
				fmt.Printf(`---------------------------------
Test Failed for message ID %s
Expecting: %v
Actual:    %v
Fail
`, msg.id, test.expected[i], msg.tags)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed for message ID %s
Expecting: %v
Actual:    %v
Pass
`, msg.id, test.expected[i], msg.tags)
			}
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
