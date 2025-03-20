/*

Tricky Slices
The append() function changes the underlying array of its parameter AND returns a new slice. This means that using append() on anything other than itself is usually a BAD idea.

// dont do this!
someSlice = append(otherSlice, element)

Take a look at these head-scratchers:

Example 1: Works As Expected
a := make([]int, 3)
fmt.Println("len of a:", len(a))
fmt.Println("cap of a:", cap(a))
// len of a: 3
// cap of a: 3

b := append(a, 4)
fmt.Println("appending 4 to b from a")
fmt.Println("b:", b)
fmt.Println("addr of b:", &b[0])
// appending 4 to b from a
// b: [0 0 0 4]
// addr of b: 0x44a0c0

c := append(a, 5)
fmt.Println("appending 5 to c from a")
fmt.Println("addr of c:", &c[0])
fmt.Println("a:", a)
fmt.Println("b:", b)
fmt.Println("c:", c)
// appending 5 to c from a
// addr of c: 0x44a180
// a: [0 0 0]
// b: [0 0 0 4]
// c: [0 0 0 5]

With slices a, b, and c, 4 and 5 seem to be appended as we would expect. We can even check the memory addresses and confirm that b and c point to different underlying arrays.

Example 2: Something Fishy
i := make([]int, 3, 8)
fmt.Println("len of i:", len(i))
fmt.Println("cap of i:", cap(i))
// len of i: 3
// cap of i: 8

j := append(i, 4)
fmt.Println("appending 4 to j from i")
fmt.Println("j:", j)
fmt.Println("addr of j:", &j[0])
// appending 4 to j from i
// j: [0 0 0 4]
// addr of j: 0x454000

g := append(i, 5)
fmt.Println("appending 5 to g from i")
fmt.Println("addr of g:", &g[0])
fmt.Println("i:", i)
fmt.Println("j:", j)
fmt.Println("g:", g)
// appending 5 to g from i
// addr of g: 0x454000
// i: [0 0 0]
// j: [0 0 0 5]
// g: [0 0 0 5]

In this example, however, when 5 is appended to g it overwrites j's fourth index because j and g point to the same underlying array. The append() function only creates a new array when there isn't any capacity left. We created i with a length of 3 and a capacity of 8, which means we can append 5 items before a new array is automatically allocated.

Again, to avoid bugs like this, you should always use the append function on the same slice the result is assigned to:

mySlice := []int{1, 2, 3}
mySlice = append(mySlice, 4)




Q.Why is 5 the final value in the last index of 'j'?

j and g point to the same underlying array so g's append overwrote j  = answer
The Go team be trollin'
Because append only works properly when the number of elements is < 10
Because 5 was appended to i to create j


Q.Why doesn't the bug regarding slices 'j' and 'g' in example 2 occur in example 1 as well?

The array's cap() is exceeded so a new underlying array is allocated = answer
The array in example 1 had a larger starting capacity
Because there are fewer elements and Go's runtime can't handle more than ~8


Q.How can you best avoid these kinds of bugs?

Don't use the append() function
Always assign the result of the append() function to a new slice
Always assign the result of the append() function back to the same slice = answer