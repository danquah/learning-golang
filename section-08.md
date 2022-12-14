# Section 8 - Control flow

Again, this is going to be a lot of basics, but there will most likely be some
small details that are special for Go, so I'll slog through it and pay attention.

A computers basic control-flows are

- Sequential - one statement at the time
- Loop - repeat a section/block of statements
- Conditional - execute a section/block based on a condition


## for

Go has `for`, looping - that it. No `while`, no `until`, no `foreach` etc. There
is only `for`.

`for` has 3 sections

- init
- condition
- post

In EBNF

```EBNF
ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
InitStmt = SimpleStmt .
PostStmt = SimpleStmt .
```

Like so

```go
for i := 0; i < 10; i++ {

}
```

[For has 2 other forms](https://go.dev/doc/effective_go#for).

If we only have a single statement, it is interpreted as the `[ Condition ]`.
This in effect gives us a while-loop.

```go
i := 0
for i < 10 {
  // Keep looping until the condition is true.
  i++
}
```

If we leave all out we get an infinite loop.

```go
for {
  // go on forever
  // unless
  if someCondition {
    break
  }
}
```

## if

The basic syntax for an if-statement is an expression (no parenthesizes) followed
by a block that will be executed if the statement is true.

```go
if someCondition {
  break
}

// Eg.
x := 42
if x > 10 {
  fmt.Println("yay")
}
```

If-statements can have an initialization statement. Any variable declared here
is scoped to the if-statement.

```go
if initialization; someCondition {
  break
}

// Eg.
if x := 42; x > 10 {
  fmt.Println("yay")
}

// This won't run.
fmt.Println(x)
```

# switch

A switch statement looks like this.

```go
switch expression {
  case expression:
    statements
  case expression:
    statements
  case expression:
    statements
}

// Eg.
x := 1
switch {
case x > 2:
  return 'nope'
case x > 0:
  return 'yup'
}
```

You can leave out the initial condition in which case it defaults to true.

```go
switch {
  case expression:
    statements
  case expression:
    statements
  case expression:
    statements
}

// Eg.
switch {
case 0 > 1:
  return 'nope'
case 1 > 0:
  return 'yup'
}
```

Go does not have automatic fall-through, so when case runs completes, the switch
terminates. Go does have a `fallthrough` keyword can be used to evaluate the
remaining conditions - but we really should not use it..

The `default` case can be used to have a case that always matches if not other
case matches.
