package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("I'm done")
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndUpper("jangwonseok")
	fmt.Println(totalLength, upperName)
}
