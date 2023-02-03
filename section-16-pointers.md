# Section 16 - Pointers

We work with pointers by using the `&` and `*` operators.

`&` gives us the address of a variable, `*` dereferences the address back to
the variable value.

In general we should just remember that everything in Go is pass by value, and
a pointer is just another value. We can use pointers in situations where passing
the actual value we want to work with gets to is a bad idea. Instead we can just
pass a references to the value and then dereference that reference back to the
variable inside the function.

The following sets up a variabel and prints out the address to it

```go
myVar := "some string"
myPointer := &myVar

fmt.Println(myPointer)
```

We can get back to the variable by dereferencing

```go
// This is basically the same as doing copyOfMyVar := myVar
copyOfMyVar := *myPointer
fmt.Println(copyOfMyVar)
```

A pointer type is denoted `*<type>`, eg. `*int`.

```go
func foo(y *int) {
  *y = *y + 1
}
```

## Method sets

Method sets is the set of methods we can access for a given type.

```go
type myType interface {
  foo() int
}

func (m myType) doStuff {

}
```

The method set of the type `myType` is now `[ doStuff ]`.

When a method attaches to a receiver, it has the option of referencing the
receiver by pointer. If it does so, it is only valid to address the method via
a pointer.

```go
type myStruct struct {
}

// This method attaches via a pointer, which tells the compiler that the method
// may want to mutate the receiver. The compiler should therefor do what it can
// to avoid situations where the receiver the method works on is a copy.
func (s *myStruct) takesPointer() {
  fmt.Printf("%T\n", s)
}

type hasPointerReceiver interface {
  takesPointer()
}

func callByPointer(thing hasPointerReceiver) {
  thing.takesPointer()
}
```

The compiler will fail on the following code.

```go
structInstance := myStruct{}
callByPointer(structInstance)
```

The compiler rejects the code to protect us from a potential bug. When `callByPointer`
is called, it is passed a _copy_ of `structInstance`, meaning that any mutation done
inside the `takesPointer` method is lost. The compiler can see that `callByPointer`
references the struct via an interface that explicitly has a receiver that wants a
pointer, and thus it refuses the call.

The compiler will refuse to call the method "by value" as any operations it makes
on its dereferenced pointer will get lost as the callByPointer function is
operating on a copy of `structInstance`. The compiler _could_ have chosen to just
add an `&` for us, but this would have been dangerous as the call would no longer
be pass by value and the semantics might have changed.

Instead we have to be explicit about the pointer:
```go
structInstance := myStruct{}
callByPointer(&structInstance)
```
