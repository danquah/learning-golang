# Section 10 - Grouping data

Collecting a bunch of data of the same type is also known as aggregating data.

When we combine several types, its called compositing. When we compose we in
effect create a new type, while aggregation is more about grouping chunks of data
of the same type together.

## Arrays

So first off, Go discourages the everyday use of Arrays. That is, use them if
you actually need them, but you it's really rare that you have to use a raw
array.

An array can contain a fixed number of elements of a type. Both the length and
element-type is a part of the arrays type. This means that it quickly becomes
very cumbersome to use arrays for anything but very fixed chunks of data.

Arrays are defined and used like so

```go
var x [5]int
x[2] = 99
fmt.Println(x)
fmt.Println(x[2])
// Outputs 5
fmt.Println(len(x))
```

## Slices

Slices is what we should use instead of arrays.

Slices lets you group values of the same type. Just like arrays, but where arrays
are static in size, slices are dynamic. They are still backed by static arrays,
but we don't have to interact directly with the arrays.
.
Slices can be defined via a composite literal. A composite literal lets us
instantiate a value that may involve multiple types.

```go
// Syntax for a slice is
// []type
// eg. []int
//
// Syntax for composite literal type{elements}
// eg. []int{1,2,3}
```

### Iterating slices

For that we use the `range` keyword.

```go
x := []int{1, 2, 3, 8}

for i, v := range x {
  fmt.Println(i, v)
}
```

### Slicing slices

You can access a slice of a slice via `:`

```go
x := []int{1, 2, 3, 8}

// Entire slices
fmt.Println(x[0:])
// Last element
fmt.Println(x[3:])
// Last element
fmt.Println(x[2:4])
```

### Appending to slices

We use `append`

```go
slice1 := []int{1,2,3}
slice1 = append(slice1, 4, 5, 6)

slice2 := int[]{7, 8, 9}

// To append an existing slices, we have to expand the previous using ...
slice1 = append(slice1, slice2 ...)
```

### Deleting for a slice

Use slicing

```go
slice1 := []int{1, 2, 3}
slice1 = append(slice1[0:1], slice1[2:]...)
fmt.Println(slice1)
```

### Making slices

A slices is backed by an underlying array. As the slice grows / and shrinks(?)
the array has to be swapped out for one that better matches the needed capacity.

We can ignore this while using the composite literal and append/delete. But, if
we're in a situation where we have a predictable growth / shrink patten, we can
use `make` to pre-create a slice with the capacity we know we'll need.

```go

// Make a slice that currently has 10 elements, but is ready to have a capacity
// of 100.
// make (type, length, capacity)
x := make([]int, 10, 100)

// We can read out the capacity using cap
fmt.Println(cap(x))
```

### Multidimensional slices

Slices are first order types, so we can have slices of slices

```go
slice1 := []int{1, 2, 3}
slice1 = append(slice1[0:1], slice1[2:]...)
fmt.Println(slice1)

mSlice1 := []string{"muggie", "moo"}
mSlice2 := []string{"foo", "bar"}

multiSlice := [][]string{mSlice1, mSlice2}
fmt.Println(multiSlice)
```

### Maps

A map is a key/value datastructure. Maps are indexed by a static index type
that provides access elements of a static element type.

The EBNF for a map

```ebnf
MapType     = "map" "[" KeyType "]" ElementType .
```

A quick way to instantiate a map is via a composite literal

```go
m := map[string]int{
  "tree": 10,
  "shrub": 4,
}
fmt.Println(m)
```

As opposed to array access, accessing an element in a map that does not exist
will give you the zero value

```go
m := map[string]int{
  "tree": 10,
  "shrub": 4,
}

// Prints 0
fmt.Println(m["mountain"])
```

We use the "comma ok" idiom to check for elements

```go
m := map[string]int{
  "tree": 10,
  "shrub": 4,
}

if v, ok := m["mountain"]; ok {
  fmt.Println("Only prints if the mountain was there ", v)
}

```

Deleting entries from a map

```go
m := map[string]int{
  "tree": 10,
  "shrub": 4,
}

delete(m, "shrub")

// Be aware this also works without any errors

delete(m, "muggie")
```
