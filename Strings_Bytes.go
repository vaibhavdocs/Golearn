package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s string
	s = "Hello, Vaibhav"
	b := []byte(s)
	c := bytes.Split(b, []byte(","))
	fmt.Printf(" % v % v\n", string(c[0]), string(c[1]))
	fmt.Printf("%q\n", bytes.Split(b, []byte(",")))
}
