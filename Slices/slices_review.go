/*

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimensions such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. If a function takes a slice argument, any changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer (we'll cover pointers later) to the underlying array. A Read function can therefore accept a slice argument rather than a pointer and a count; the length within the slice sets an upper limit of how much data to read. Here is the signature of the Read() method of the File type in package os:

Referenced from Effective Go

func (f *File) Read(buf []byte) (n int, err error)



Q.Which references the other?

Slices reference arrays = answer
Arrays reference slices



Q.Multiple slices can point to the same array

True = answer
False



Q.A function that only has access to a slice can modify the underlying array

True = answer
False