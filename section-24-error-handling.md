# Section 23: Error handling

Gos approach to handling errors and potential errors, is to handle the issue
as close as possible to where it occurs. This means handling an error the moment
you realize it happens, but also guarding for potential errors before they happen.

Errors in Go are described by a simple interface

```go
type error interface {
  Error() string
}
```

## Error checking

A function that returns other business data can return an error by using
Gos multiple return values feature. Should always check an returned error.

```go
err = someOtherFunction()
if err != nil {
  fmt.Println(err)
  os.Exit(1)
}

data, err := someFunction()
if err != nil {
  fmt.Println(err)
  os.Exit(1)
}
```

## Handling errors before the occur

Potential errors should also be handled as quickly as possible. The closing of
a file we've just opened can for instance be handled via defer:

```go
  f, err := os.Create(filePath)
  if err != nil {
    log.Fatal("Could not create file ", filePath, " :", err)
  }
  // Queue up the close right away.
  defer f.Close()
  f.WriteString(token)
```

## Handling errors, logging

The `log` package gives us some nice ways of logging errors right out of the box.

Where `fmt.Print*` logs to stdout, `log.Print*` and others has the option of
logging to other streams such as a file. Log also gives us timestamp.

```go
log.Println("hello")
// Outputs 2023/01/31 07:20:14 Hello
```

The `log` has functions for logging and then terminating

```go
// Shut down with a non-zero exit code, no deferred functions will be called.
log.Fatalln("Hello")
// Print and panic - deferred functions will be called.
log.Panicln("Hello")
```

## Defer

Some more details on defer

1. A deferred functions arguments are evaluated when it is called
2. Deferred functions calls are executed last in first out - like a stack
3. Deferred functions may read and assign to the returning functions named return
   values

## Panic / Recover

Panic / recover are control-flow constructs. When panic is called go will start
unwinding the call-stack, calling any deferred functions fifo.

A deferred function can recover. That is, call `recover()` and have normal execution
resume from that point on.

```go
func foo() {
  log.Println("I'm foo")

  // Are we panicking?
  if i := recover(); i != nil {
    log.Println("Recovering from ", i)
  }
}
```

## Adding information to errors

The `error` package has a constructor for creating errors

```go
func haveProblems() (int, error) {
  return 0, errors.New("i'm bad at this")
}
```

If we just want to some simple formatting of the error message we can use a
convenience function from `fmt`

```go
fmt.Errorf("Something bad happend with %v", variable)
```

We can create our own errors by implementing the `error` interface.

```go

// The interface looks like this:

type error interface {
  Error() string
}

// Setup our own struct with a custom field for our error to use.

type thingy struct {
  value string
}

// Then implement the error interface
func (t thingy) Error() string {
  return fmt.Sprint("Custom error string for ", t.value)
}

// And we can now use the error.
func haveProblems() (int, error) {
  return 0, thingy{"stuff happened"}
}
```
