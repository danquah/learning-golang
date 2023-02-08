# Section 28: Testing

A test must be in a file that ends in `_test.go` - the test should be in the same
package as the code being tested.

The test functions must start with `Test` and then proceede with something
meaningfull. If we're testing a specific function named eg. `FunctionName`, it is
best practice to name the function `TestFunctionName`.

A test uses the `testing` package, which gives us a bunch of useful functions for
handling the testing flow.

```go
package main

import "testing"

func TestMyFunction(t *testing.T) {
  expected := "value"
  result, err := myFunction

  if err != nil {
    t.Error("Something bad happend ", err)
  }

  if result != expected {
    t.Error("Unexpected return ", result, " expected ", expected)
  }
}
```

## Test examples

We can embed examples into our documentation


