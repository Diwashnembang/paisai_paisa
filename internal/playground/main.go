package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world nepal"
	fmt.Println(strings.TrimPrefix(s, "hello world"))
}
