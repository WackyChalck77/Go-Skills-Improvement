package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString("123")
	buffer.WriteString("world!")
	result := buffer.String()
	fmt.Println(result)

	var builder strings.Builder
	builder.WriteString("Hello, ")
	builder.WriteString("world!")
	result2 := builder.String()
	fmt.Println(result2)
}
