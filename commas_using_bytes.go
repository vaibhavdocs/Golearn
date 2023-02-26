package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "-123.2"
	comma(s)
}

//Non-recursive function for the comma using bytes.Buffer
func comma(s string) {
	var b bytes.Buffer
	if len(s) > 3 {
		b.Write([]byte(s[:3]))
		b.WriteByte(',')
		b.Write([]byte(s[3:]))

	}
	fmt.Printf(b.String())
}

//
