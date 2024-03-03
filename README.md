# go-result-and-option

This Golang package implements two generic types `Result[T]` and `Option[T]`, they are inspired by Rust's enum `Result<T,E>` and `Option<T>` in it's standard library.

## Installation

```shell
go get -u github.com/yuanzicheng/go-result-and-option
```

## Examples

```go
package main

import (
	"fmt"

	"github.com/yuanzicheng/go-result-and-option/option"
	"github.com/yuanzicheng/go-result-and-option/result"
)

func main() {
	i := 12345
	x := 123

	{
		res := result.Ok(&i)
		v := res.UnwrapOr(&x)
		fmt.Println(*v) // 12345
	}

	{
		opt := option.None[int]()
		v := opt.UnwrapOr(&x)
		fmt.Println(*v) // 123
	}
}
```
