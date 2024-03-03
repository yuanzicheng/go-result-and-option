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
