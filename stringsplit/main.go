package main

import (
	"fmt"
	"strings"
)

type TestCase struct {
	s   string
	sep string
}

func main() {
	cases := []TestCase{
		{"1-2-3", "-"}, // 3 [1, 2, 3]
		{"123", "-"},   // 1 [123]
		{"", "-"},      // 1 []
		{"123", ""},    // 3 [1, 2, 3]
		{"", ""},       // 0 []
	}

	for _, c := range cases {
		fmt.Printf("\ncase: %#v\n", c)
		split := strings.Split(c.s, c.sep)
		fmt.Printf("\tresult slice: len = %d, value = %#v\n", len(split), split)
	}
}
