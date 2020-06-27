// 3.10 - Write a non-recursive version of `comma`, using `bytes.Buffer` instead
// of string concatenation.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	// Check if sign, if there is then set start, and write the sign to buf
	start := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		start = 1
	}

	// Check if decimal, if there is not then set end to length of input string
	end := strings.Index(s, ".")
	if end == -1 {
		end = len(s)
	}

	// Create substring slice from start, and end for evaluation of commas
	t := s[start:end]

	// Determine index of first comma within substring
	n := len(t) % 3
	if n == 0 {
		n = 3
	}

	// Write into buffer the substring up to first comma then iterate by 3
	// to insert more comma's into buffer if substring is long enough
	buf.WriteString(t[:n])
	for i := n; i < len(t); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(t[i : i+3])
	}

	// If substring length is less than input string length then write
	// decimal point and fraction into buffer
	if len(t) < len(s) {
		buf.WriteByte('.')
		buf.WriteString(s[end+1:])
	}

	return buf.String()
}

func main() {
	s := []string{
		"12345",
		"23",
		"0",
		"1234567890",
		"+123.12",
		"-123456.789",
		"+1234567890.123456",
	}
	for _, e := range s {
		fmt.Printf("comma(%s)=%s\n", e, comma(e))
	}
}
