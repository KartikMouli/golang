/*

Memory Layout
In Go, structs sit in memory in a contiguous block, with fields placed one after another as defined in the struct. For example this struct:

type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

Looks like this in memory:

struct layout

Field ordering... Matters?
the order of fields in a struct can have a big impact on memory usage. This is the same struct as above, but poorly designed:

type stats struct {
	NumPosts uint8
	Reach    uint16
	NumLikes uint8
}

It looks like this in memory:

struct layout

Notice that Go has "aligned" the fields, meaning that it has added some padding (wasted space) to make up for the size difference between the uint16 and uint8 types. It's done for execution speed, but it can lead to increased memory usage.

Should I Panic?
To be honest, you should not stress about memory layout. However, if you have a specific reason to be concerned about memory usage, aligning the fields by size (largest to smallest) can help. You can also use the reflect package to debug the memory layout of a struct:

typ := reflect.TypeOf(stats{})
fmt.Printf("Struct is %d bytes\n", typ.Size())

Real Story
I once had a server in production that held a lot of structs in memory. Like hundreds of thousands in a list. When I re-ordered the fields in the struct, the memory usage of the program dropped by over 2 gigabytes! It was a huge performance win.

Assignment
Our over-engineering boss is at it again. He's heard about memory layout and wants to squeeze every last byte out of our structs.

Run the tests to see the current size of the structs, then update the struct definitions to minimize memory usage.


*/

package main

import (
	"fmt"
	"reflect"
	"testing"
)

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

func Test(t *testing.T) {
	type testCase struct {
		name     string
		expected uintptr
	}

	runCases := []testCase{
		{"contact", uintptr(24)},
		{"perms", uintptr(16)},
	}

	submitCases := append(runCases, []testCase{}...)

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		var typ reflect.Type
		if test.name == "contact" {
			typ = reflect.TypeOf(contact{})
		} else if test.name == "perms" {
			typ = reflect.TypeOf(perms{})
		}

		size := typ.Size()

		if size != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v bytes
Actual:     %v bytes
Fail
`, test.name, test.expected, size)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v bytes
Actual:     %v bytes
Pass
`, test.name, test.expected, size)
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
	mockTests := &testing.T{}
	Test(mockTests)
}
