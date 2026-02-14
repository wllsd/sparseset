# Sparse Set

A high-performance, generic Sparse Set implementation for Go.

## Overview

The Sparse Set is a data structure designed for performance-critical scenarios, such as Entity Component Systems (ECS). It allows for efficient iteration over packed data while maintaining a mapping to sparse indices, avoiding the overhead of hash maps or the fragmentation of standard arrays.

## Usage

```go
package main

import (
    "fmt"
    "github.com/wllsd/sparseset"
)

func main() {
    set := sparseset.New[string](1000)

    set.Add(10, "Foo")
    set.Add(20, "Bar")

    if set.Has(10) {
        fmt.Println("ID 10 exists")
    }

    for id, value := range set.All() {
        fmt.Printf("ID: %d, Value: %s\n", id, value)
    }

    set.Remove(10)
}
```

