package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

const (
	message = "Hello, OTUS!"
)

func main() {
	fmt.Println(stringutil.Reverse(message))
}
