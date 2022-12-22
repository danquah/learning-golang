# Section 12 structs

Declared like so

```go
type person struct {
  first string,
  last string,
}

We can then init a struct with a composite literal

```go

dude := person{
  first: "the",
  last: "dude",
}

```

Fields of a struct are accessed via dot notation

```go
fmt.Println(dude.first)
```

## Embedding

You can embed a struct into another struct, either with an anonymous field or
a named

```go
type person struct {
  first string
  last string
}


// Anonymous, the name of the field will be set to the types name
type importantPerson struct {
  person
  role string
}

// Named
type twins struct {
  person1 person
  person2 person
  role    string
}
```

The fields of the inner type gets promoted to the outer

```go

santa := importantPerson{
		person: person{
			first: "Santa",
			last:  "Claus",
		},
		role: "Distribute presents",
	}

fmt.Println(santa.first)

// But you can also specify the field relative to the inner type
fmt.Println(santa.person.first)
```

## Anonymous structs

We can combine the definition and initialization of a struct for a one-time use
anonymous struct.

```go

somePerson := struct {
  first string
  last  string
}{
  first: "James",
  last:  "Bond",
}
fmt.Println(somePerson)

```
