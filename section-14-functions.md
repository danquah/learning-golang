# Section 14 - functions

A function has a receiver, identifier, parameters, return types and a body.

```go
func (r reciver) identifier(parameters) (returntype) {
  body
}

// eg
func (yadda) myFuncName(s string) bool {
  if s == "moo" {
    return false
  }

  return true
}
```

## Return

A function can return multiple values.
```go
// Multi return
func (yadda) multiReturn(s string) (bool, string) {
  if s == "moo" {
    return false, "moo"
  }

  return true, "foo"
}

a, b := multiReturn("moo")

```

## Variadic parameters

A function can take a variable number of parameter, this is call a variadic parameter.

The variadic parameter must be the last parameter.

```go

func multiParameter(something int, s ...string) {
  for i, v := range s {
    fmt.Println(something, i, v)
  }
}
```

The variadic parameter will be of type []T. It is possible to pass in an existing
slice as an argument by following the parameter name with `...`

```go
func variadicFunction(s ...string) {
  fmt.Printf("Slice has length: %v, capacity: %v, data: %v\n", len(s), cap(s), s)

  for i, v := range s {
    fmt.Println(i, v)
  }
}
maSlice := []string{"boom", "boom", "boom"}
variadicFunction(maSlice...)
```

## Defer

We can defer the call of a function to when the scope is exited

```go
func foo() {
  fmt.Println("Foo")
}

func bar() {
  fmt.Println("Bar")
}

func useDefer() {
  defer foo()
  bar()
}
// Outputs:
// Bar
// Foo

```

## Methods

As pr the [spec](https://go.dev/ref/spec#Method_declarations)

> A method is a function with a receive

A method is declared like so

```go
func (r receiver) identifier(parameters) (returnType) {
  body
}
```

Eg.


```go
// Declare a type we can attach to.
type PersonName struct {
  first string
  last  string
}

// Then attach.
func (p PersonName) greet(say string) {
  fmt.Printf("%v %v %v", say, p.first, p.last)
}

// Can then be used like so
type PersonName struct {
  first string
  last  string
}

func (p PersonName) greet(say string) {
  fmt.Printf("%v %v %v", say, p.first, p.last)
}

func main() {
  person1 := PersonName{first: "Some", last: "Person"}
  person1.greet("Hello")
}

```

## Interfaces

An interface declares the pattern a type must match in other to be of a type.

This means that a value can be of more than one type.

An interface is declared like so

```go
type Speaker interface {
  speak()
}
```

Which declares any type that has `speak()` attached to be of type `Speaker`.
We can then let different types implement the interface and write more generic
code:

```go
// Declare different types
type Animal struct {
  name string
}

type Human struct {
  first string
  last  string
}

// Then attach methods to them that makes them implement Speaker
func (a Animal) speak(message string) {
  fmt.Printf("%v says: %v\n", a.name, message)
}

func (h Human) speak(message string) {
  fmt.Printf("%v %v says: %v\n", h.first, h.last, message)
}

// Setup a method that uses a speaker
func react(speaker Speaker, message string) {
  speaker.speak(message)
}
```

## The empty interface

The empty interface can be used in situations where you want to accept any type.

```go
func veryFlexible(canBeAnything interface{}) {
  ...
}
```

## Assertion and type switch

We may want to go from the interface back to the concrete type. This can be done
by asserting the type.

```go
// Type assertion
fmt.Println(speaker.(Human).first)
// The following would panic
// fmt.Println(speaker.(Animal).first)
```

Assertions will panic if the value is not of the asserted type. To dynamically
handle the type of a value, we can switch on the type instead.

```go
switch speaker := speaker.(type) {
case Animal:
  fmt.Println("The speaker is an animal named ", speaker.name)
case Human:
  fmt.Println("The speaker is a human named ", speaker.first)
default:
  fmt.Println("No idea which typer of speaker this is")
}
```

## Anonymous functions

Functions can be anonymous.

We can run the function right away, these are called anonymous self-executing
functions

```go
// Or we can run the function right away
func() {
  fmt.Println("I'll run right away")
}()
```

Or we can store the function as a value

```go
myFunc := func() {
  fmt.Println("Im anonymous")
}

// Run the function.
myFunc()
myFunc()

// We can also pass the functions around
// This is a function that returns func that takes an int aka func(int)
makeANumberCaller := func() func(int){
  return func(num int){
    fmt.Printf("The number is %d", num)
  }
}

numberCaller := makeANumberCaller()
numberCaller(42)
```

## Code blocks

We can at any time create an arbitrary code block to create a scope

```go
{
  anEnclosedVar := "secret stuff"
  fmt.Println(anEnclosedVar)
}

// This would not compile.
fmt.Println(anEnclosedVar)

```

## Closures

We can let a returned function carry a value from the scope it was defined in:

```go
func returnEnclosed() func() {
  inTheEnclosedScope := "Some value"

  return func() {
    fmt.Printf("The enclosed value is %v", inTheEnclosedScope)
  }
}
```

This way we can keep values internal to a function, but still let external code
interact with a function that has access to the value.


