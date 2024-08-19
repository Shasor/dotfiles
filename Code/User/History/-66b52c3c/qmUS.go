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
