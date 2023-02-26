package main

import (
	"fmt"
	"sort"
)

func main() {
	isAnagram("anagram", "nagagram")
}

func isAnagram(s1 string, s2 string) {
	r1 := []rune(s1)
	r2 := []rune(s2)
	// sort.Ints(r1)
	fmt.Println(sort.Ints(r1))

	for i := range r1 {
		if (64 < r1[i] && r1[i] < 91) || (96 < r1[i] && r1[i] < 123) {
			for j := i; j < len(r2); j++ {
				if (64 < r2[i] && r2[i] < 91) || (96 < r2[i] && r2[i] < 123) {
					if r1[i]-r2[j] == 0 {
						count++
						fmt.Println(r1[i], r2[j])
						fmt.Println("something")
						break
					}
				}

			}
		}

	}
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(count)

}
