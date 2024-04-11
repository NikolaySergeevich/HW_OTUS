package main

import (
	"fmt"                                                    //nolint
	"github.com/abelit/abelit-code-go/src/gobase/stringutil" //nolint:depguard
)

func main() {
	hello := "Hello, Otus!"
	fmt.Println(stringutil.Reverse(hello))
}
