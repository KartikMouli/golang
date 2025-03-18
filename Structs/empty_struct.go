/*

Empty Struct
Empty structs are used in Go as a unary value.


// anonymous empty struct type
empty := struct{}{}

// named empty struct type
type emptyStruct struct{}
empty := emptyStruct{}

The cool thing about empty structs is that they're the smallest possible type in Go: they take up zero bytes of memory.

memory usage

Later in this course, you'll see how and when they're used: it's surprisingly often! Mostly with maps and channels.

*/


Q.Why does the empty anonymous struct have two pairs of braces? 'struct{}{}'



It doesn't, it's a syntax error


Because the Go developers like to flex their 200 WPM typing speed


'struct{}' is the type (empty struct) and '{}' is the value (empty struct literal)  = answer


Become a member for quiz access



Q.Which is ordered from least -> most memory usage?



uint16, bool, int64, struct{}


struct{}, uint16, bool, int64


struct{}, bool, uint16, int64 = answer


bool, struct{}, uint16, int64

