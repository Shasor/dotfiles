package main

func Max(s []int) int {
	var r int
	for i := 1; i < len(s); i++ {
		a, b := s[i-1], s[i]
		if a > b {
			r = a
		}
	}
	return r
}