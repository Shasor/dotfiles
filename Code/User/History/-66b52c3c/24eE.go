package main

import (
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]
	
	idx1, idx2 := 0, 0

	for idx1 < len(s1) && idx2 < len(s2) {
		if s1[idx1] == s2[idx2] {
			idx1++
		}
		idx2++
	}

	if idx1 == len(s1) {
		for _, c := range s1 {
			z01.PrintRune(c)
		} 
		z01.PrintRune('\n')
	}
}
