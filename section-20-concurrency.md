# Section 20: Concurrency

One of the main goals of Go was to make concurrency easy. The Go programming language
was started in 2007 - one year after Intel released their first dual-core processor.

Taking advantage of multi-core processors, requires the code to be concurrent.
That is, code that *can* run in parallel. Concurrency is a design pattern that
allows you to take advantage of multiple cpus/cores, but does not rely on it.

Go uses "Go rutines" to handle concurrency. We launch a rutine using the `go` keyword:

```go
func foo(){
  ...
}

func main() {
  go foo()
}
```

## WaitGroups

When we have concurrency, we need coordination. WaitGroups is one approach that
allows us to wait for things to complete.

```go
var wg sync.WaitGroup

func foo(){
  ...
  wg.Done()
}

func main() {
  // Set the number of running things
  wg.Add(1)
  // Launch a Go rutine
  go foo()
  // do some stuff
  // Then wait for foo
  wg.Wait()
}
```

## Mutex

Go rutines can potentially access a shared variable. To coordinate the access
to a variable, we can use a mutex.

```go
var mu sync.Mutex

go func(){
  mu.Lock()
  // Access shared variable
  mu.Unlock()
}()
```

Go also has mutexes for more detailed locking of read/write:

```go
var mu sync.RWMutex

go func(){
  mu.RLock()
  // Read shared variable
  mu.RUnlock()
  mu.Lock()
  // Write to the shared variable
  mu.Unlock()
}()
```

## Atomic

The `sync.atomic` package has a number of tools for more detailed concurrency work.

Instead of handling mutexes ourselves we can do this:

```go
var counter int64

// In some loop that starts a lot of rutines.
go func(){
  atomic.AddInt64(&counter, 1)
  runtime.Gosched()
  fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
}()
```
