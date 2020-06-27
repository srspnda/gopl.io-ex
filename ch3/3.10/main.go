// 3.10 - Write a non-recursive version of `comma`, using `bytes.Buffer` instead
// of string concatenation.
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	n := len(s) % 3
	if n == 0 {
		n = 3
	}

	buf.WriteString(s[:n])

	for i := n; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

func main() {
	s := []string{
		"12345",
		"23",
		"0",
		"1234567890",
	}

	for _, e := range s {
		fmt.Println("comma(", e, ") = ", comma(e))
	}
}
