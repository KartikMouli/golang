/*

Process Notifications
Textio now has a system to process different types of notifications: direct messages, group messages, and system alerts. Each notification type has a unique way of calculating its importance score based on user interaction and content.

Assignment
Implement the importance methods for each message type. They should return the importance score for each message type.
For a directMessage the importance score is based on if the message isUrgent or not. If it is urgent, the importance score is 50 otherwise the importance score is equal to the DM's priorityLevel.
For a groupMessage the importance score is based on the message's priorityLevel
All systemAlert messages should return a 100 importance score.
Complete the processNotification function. It should identify the type and return different values for each type
For a directMessage, return the sender's username and importance score.
For a groupMessage, return the group's name and the importance score.
For a systemAlert, return the alert code and the importance score.
If the notification does not match any known type, return an empty string and a score of 0.

*/

package main

import (
	"fmt"
)

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int {
	if dm.isUrgent == true {
		return 50
	}

	return dm.priorityLevel
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

func (sa systemAlert) importance() int {
	return 100
}

// ?

func processNotification(n notification) (string, int) {
	// ?
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, v.importance()
	case groupMessage:
		return v.groupName, v.importance()
	case systemAlert:
		return v.alertCode, v.importance()
	default:
		return "", 0
	}
}


func main() {
	testCases := []struct {
		notification       notification
		expectedID         string
		expectedImportance int
	}{
		{
			directMessage{senderUsername: "Kaladin", messageContent: "Life before death", priorityLevel: 10, isUrgent: true},
			"Kaladin",
			50,
		},
		{
			groupMessage{groupName: "Bridge 4", messageContent: "Soups ready!", priorityLevel: 2},
			"Bridge 4",
			2,
		},
		{
			systemAlert{alertCode: "ALERT001", messageContent: "THIS IS NOT A TEST HIGH STORM COMING SOON"},
			"ALERT001",
			100,
		},
		{
			directMessage{senderUsername: "Shallan", messageContent: "I am that I am.", priorityLevel: 5, isUrgent: false},
			"Shallan",
			5,
		},
		{
			groupMessage{groupName: "Knights Radiant", messageContent: "For the greater good.", priorityLevel: 10},
			"Knights Radiant",
			10,
		},
	}

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		id, importance := processNotification(test.notification)
		if id == test.expectedID && importance == test.expectedImportance {
			passCount++
			fmt.Printf("Test %d Passed: Expected (%s, %d), Got (%s, %d)\n", i+1, test.expectedID, test.expectedImportance, id, importance)
		} else {
			failCount++
			fmt.Printf("Test %d Failed: Expected (%s, %d), Got (%s, %d)\n", i+1, test.expectedID, test.expectedImportance, id, importance)
		}
	}

	fmt.Printf("\nSummary: %d passed, %d failed\n", passCount, failCount)
}
