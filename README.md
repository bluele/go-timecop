# go-timecop

go-timecop provides "time travel", "time freezing" capabilities for testing.
This is inspired by RubyGems [timecop](https://github.com/travisjeffery/timecop).

# Getting started

## Installation

```
$ go get -u github.com/bluele/go-timecop
```

## Example

```go
// examples/examples.go
package main

import (
  "fmt"
  "github.com/bluele/go-timecop"
)

func main() {
  t := timecop.Now()
  fmt.Printf("current: %v\n", t)
  fmt.Println("Dive into the future!")
  timecop.Travel(t.AddDate(1, 0, 0))

  for i := 0; i < 3; i++ {
    fmt.Printf("future: %v\n", timecop.Now())
  }

  fmt.Println("Return to the current.")
  timecop.Return()

  fmt.Printf("current: %v\n", timecop.Now())
}
```

```
$ go run examples/examples.go
current: 2015-05-25 09:48:16.413680379 +0900 JST
Dive into the future!
future: 2016-05-25 09:48:16.413681008 +0900 JST
future: 2016-05-25 09:48:16.41369027 +0900 JST
future: 2016-05-25 09:48:16.413697228 +0900 JST
Return to the current.
current: 2015-05-25 09:48:16.414099205 +0900 JST
```

## Difference between go-timecop and [timecop](https://github.com/travisjeffery/timecop).

Ruby can be integrated with "time module" using metaprogramming, but the same behavior is not possible using golang.

Therefore, it is necessary to replace the `time.Now` and` time.Since` with function `timecop.Now` and` timecop.Since` which is provided by this library.

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>
