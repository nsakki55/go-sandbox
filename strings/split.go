package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "sampleTest"
	splitText := strings.Split(text, ",")[1]
	fmt.Printf(splitText)
}
