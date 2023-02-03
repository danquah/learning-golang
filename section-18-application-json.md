# Section 16 - Writing an application

## Documentation

<https://pkg.go.dev/>

## Error handling

We will go more in to details in a later section, but for now it is worth noting
that error-handling in Go focuses on handling the error at the point where it
is generated. Errors are returned as values from functions / methods, and the
caller is expected to handle the error right there and then. This is different
from languages such as Java and C# where Exceptions are handled at the end of
a block.

```go
b, err := json.Marshall(group)

if err != nil {
  fmt.Println("Error: ", err)
}
```

## An application example - handling json

We're doing some rendering and parsing of JSON.

Go's json package lets us take a struct, and marshal it into a json string.
The `Marshal` method reads the public fields of a struct, and generates a matching
json object:

```go

type person struct {
  name string
  age int
}

myPerson := person{"Ole", 42}
bs, err := json.marshal(myPerson)
```

`Unmarshal` lets us move in the opposite direction

```go
// We have bs from before
var myPerson person;
err := json.unmarshal(bs, &myPerson)
// myPerson should now be initialized. with our data
```

In the (usual) situation where the properties in our json object does not match
up perfectly with the fields in our struct, we can use tags to instruct the
unmarshaling process what to do.

```go
  type LowerCasedPerson struct {
    First string `json:"first"`
    Last  string `json:"last"`
    Age   int    `json:"age"`
  }
  jsonString := `[{"first":"Jens","last":"Peter","Age":24}]`
  var lowPeople []LowerCasedPerson
  err = json.Unmarshal([]byte(jsonString), &lowPeople)
```

## `Writer` / `Reader` interfaces

Both of these interfaces are good examples of simple interfaces that are used
throughout the standard library.

The [`Writer`](https://pkg.go.dev/io#Writer) interface looks like this:

```go
type Writer interface {
  Write(p []byte) (n int, err error)
}
```

Which means that eg. the `os.File.Write` method that has the following
signature:

```go
func (f *File) Write(b []byte) (n int, err error)
```

Which means that any `File` compatible with `Writer`. All the same goes for `Reader`.

Another example of a writer is `os.Stdout`. We could obviously just use `fmt.Print*`
but lets have a go.

```go
fmt.Println("The boring way")
fmt.Fprint(os.Stdout, "Much more 1337\n")
```

## Sorting

The `sort` package has a bunch of convinient tools for letting us sort datastructures
such as `Ints` and `Strings` for sorting slices of ints and strings. It also lets
us do custom sorting, which serves as a good example of how interfaces are used
in Go for composing functionality.

Let's say we have a slice of people:

```go
type person struct {
  first string
  last string
  age int
}

// declare some people in p1-p4
people := []person{p1, p2, p3, p4}
```

We first declare an interface of the same type as the slice we want to sort:

```go
type ByAge []Person
```

This interface lets us do two important things

1. We have a place to attach functions that can then fulfill another inteface
2. It lets us keep the sorting functionality separate from the datastructure.
   that is, we don't have to attach these functions directly to our slices

To do the sorting, our `ByAge` must implement `Sort.Interface`:

```go
type Interface interface {
  Len() int
  Less(i, j int) bool
  Swap(i, j int)
}
```

Like so:

```go
func (a ByAge) Len() int {
  return len(a)
}

func (a ByAge) Less(i, j int) bool {
  return a[i].age < a[j].age
}

func (a ByAge) Swap(i, j int) {
  a[i], a[j] = a[j], a[i]
}
```

And we can now do the sorting by converting the slice to our sort interface and
passing it to `sort.Sort`:

```go
sort.Sort(ByName(people))
```
