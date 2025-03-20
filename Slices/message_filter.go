/*

Message Filter
Textio is introducing a feature that allows users to filter their messages based on specific criteria. For this feature, messages are categorized into three types: TextMessage, MediaMessage, and LinkMessage. Users can filter their messages to view only the types they are interested in.

Assignment
Your task is to implement a function that filters a slice of messages based on the message type.

Complete the filterMessages function. It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link"). It should return a new slice of Message interfaces containing only messages of the specified type.

*/

package main

import (
	"fmt"
)

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	// ?
	ans := []Message{}

	for _,msg := range messages{
		if msg.Type() == filterType{
			ans = append(ans, msg)
		}
	}

	return ans
}

func main() {
    messages := []Message{
        TextMessage{"Alice", "Hello, World!"},
        MediaMessage{"Bob", "image", "A beautiful sunset"},
        LinkMessage{"Charlie", "http://example.com", "Example Domain"},
        TextMessage{"Dave", "Another text message"},
        MediaMessage{"Eve", "video", "Cute cat video"},
        LinkMessage{"Frank", "https://boot.dev", "Learn Coding Online"},
    }

    testCases := []struct {
        filterType    string
        expectedCount int
    }{
        {"text", 2},
        {"media", 2},
        {"link", 2},
    }

    passCount, failCount := 0, 0
    for i, test := range testCases {
        filtered := filterMessages(messages, test.filterType)
        if len(filtered) != test.expectedCount {
            fmt.Printf("TestCase %d Failed: Expected %d, Got %d\n", i+1, test.expectedCount, len(filtered))
            failCount++
        } else {
            fmt.Printf("TestCase %d Passed\n", i+1)
            passCount++
        }
    }

    fmt.Printf("Summary: %d Passed, %d Failed\n", passCount, failCount)
}