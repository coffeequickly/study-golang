package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, _ := lenAndUpper("jangwonseok")
	fmt.Println(totalLength)
}
