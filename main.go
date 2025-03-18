package main

import (
	"fmt"

	"github.com/ajha-cs/go-pkg/gopkg"
)

func main() {
	message := gopkg.Greet("World")
	fmt.Println(message)
}
