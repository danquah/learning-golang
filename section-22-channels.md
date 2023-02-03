# Section 21: Channels

Channels lets us coordinate work in concurrent code.

The syntax for creating a channel, putting an element on it and reading the value off it is

```go
c := make(chan int)
c <- 42
fmt.Println(<-c)
```

Channels blocks. So, we should make sure a channel has capacity for the data we
want to put on it, unless we're coordinating via capacity.

The same goes for reading. If a channel is empty and we read from it, we block
until there is data.

You can buffer a channel by specifying a capacity on creation

```go
c := make(chan int, 3)
```

We can close a channel after which reads will return the zero-value of the type.
A multivalued assignment returns both the value and whether the channel was
closed and can be used to detected closed channels and disregard the value.

```go
  // Setup an unbuffered channel
  c := make(chan int)

  // Read data off the channel
  go func() {
    fmt.Println(<-c)
  }()

  // Send the value to be read
  c <- 42

  close(c)

  // Reads of a closed channel returns immediately so we can safely do it from
  // main
  v, ok := <-c

  // Outputs 0 false
  fmt.Println(v, ok)

````

## Directional channels

We have three possible configurations of directions for channels.

```go
// Unidirectional
chanUni := make(chan int)
// This we can only write to
chanRead := make(chan<- int)
// This we can only Read from
chanSend := make(<-chan int)
```

## Ranging

We can range over a channel, in which case closing the channel usually becomes
important.

```go
c := make(chan int)

// This will push a bunch of values
go func(){
  for i := 0; i < 100; i++ {
    c <- i
  }
  // It is important to close the channel, otherwise the range loop won't stop.
}()

// We read data of the channel with range
for v := range(c) {
  fmt.Println(v)
}

```

## Select

Select is like a switch-statement for channels. A select picks from the first
available channel:

```go
    select {
    case v := <-chan1:
      fmt.Println("Value: ", v)
    case v := <-chan2:
      fmt.Println("Value: ", v)
    }
```

## Fan in / out

Fanning in and out is about multiplexing and de-multiplexing concurrent streams.

## Context

Context handles request-scoped data that all rutines that handles the request
needs access to. A context is even more importantly used to signal all the
actors that are processing a request.


