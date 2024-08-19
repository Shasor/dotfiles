// wdmatch
// Instructions
// Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string.

// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.

// If the number of arguments is different from 2, the program displays nothing.

// Usage
// $ go run . 123 123
// 123
// $ go run . faya fgvvfdxcacpolhyghbreda
// faya
// $ go run . faya fgvvfdxcacpolhyghbred
// $ go run . error rrerrrfiiljdfxjyuifrrvcoojh
// $ go run . "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
// quarante deux
// $ go run .
// $

// Seth Solution 1
package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    // Check if the number of arguments is not equal to 2
    if len(os.Args) != 3 {
        return // Exit the program if the number of arguments is not correct
    }

    // Extract the two strings from command line arguments
    str1, str2 := os.Args[1], os.Args[2]

    // Pointers to track the position in each string
    ptr1, ptr2 := 0, 0

    // Iterate over the characters of str2
    for ptr2 < len(str2) {
        // If the current character in str2 matches the current character in str1
        if ptr1 < len(str1) && str1[ptr1] == str2[ptr2] {
            ptr1++ // Move to the next character in str1
        }
        ptr2++ // Move to the next character in str2
    }

    // If ptr1 reached the end of str1, it means all characters of str1 were found in str2
    if ptr1 == len(str1) {
        // Print the string followed by a newline
		for _, v := range str1 {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
    }
}

// Seth solution 2

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