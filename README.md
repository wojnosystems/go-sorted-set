# Overview

Provides a sorted set for the common golang primitive types. Provides a generator template for any you may wish to generate.

# How to use it

## Install

`go get github.com/wojnosystem/go-sorted-set`

## In code

```go
package main

import (
    "fmt"
	"github.com/wojnosystems/go-sorted-set"
)

func main() {
  sortedStrings := sorted_set.NewString("z", "y", "x").Sort()
  for _, s := range sortedStrings {
    fmt.Println(s)
  }
}
```

Prints:

```
x
y
z
```

NewString is the builder. You can pass in 0 or more strings to sort. You can also call "Add()" on the builder as well to add in any additional strings to sort. To create a sorted array, call "Sort()".

If you have a sorted_set.String, you know that each item is unique and is in order. You can run binary searches on this set.

# Other types

All of the other Go-Lang common types are supported (it makes no sense to have a sorted set of bools). Just use them by replacing "NewString" with the type you want, such as: `NewUint()`

# Motivation

I originally wanted this for a library that compared input values with an expected set of values. I wanted to be sure that there were no duplicates and that they were sorted to speed up searches. Originally, I had thrown errors if this structure was out of order. But checking and returning errors on each use is a hassle, so this structure has become the state I wanted those values to be in.

# FAQ

## Yea, but, I can just change the values in the arrays

Yup. This isn't a security enforcement system. But a way of telling developers that the array that they're working with is intended to be sorted at all times and never contain duplicates.

## What about set operators?

Sure. Open a PR and we'll add them in.
