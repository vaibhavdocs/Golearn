package main

import "fmt"

func main() {
	s := "🌹"
	fmt.Printf("% x\n", s)

	r := []rune(s)
	fmt.Printf("%x\n", r)

	fmt.Println(string(12323))
}
