package main

import (
	"fmt"
)

// check with go run -race main.go
func main() {
	i := 0
	go func() { i++ }()
	fmt.Println(i)
}
