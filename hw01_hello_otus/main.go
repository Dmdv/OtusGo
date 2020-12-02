package hw01_hello_otus

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

const (
	Message = "Hello, OTUS!"
)

func main() {
	fmt.Println(stringutil.Reverse(Message))
}
