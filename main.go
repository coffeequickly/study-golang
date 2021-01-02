package main

import (
	"fmt"
)

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("test", "test2", "test3")
}
