package main 

import (
	"github.com/01-edu/z01"
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

//Ferdinand Alt

package main

import (
	"fmt"
	"os"
)

func WordMatch(str1, str2 string) bool {
	if len(str1) == 0 || len(str2) == 0 {
		return false
	}

	i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}

	return i == len(str1)
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1, str2 := os.Args[1], os.Args[2]

	if WordMatch(str1, str2) {
		fmt.Println(str1)
	}
}