package main

func Max(s []int) int {
	var r int
	for i := 1; i < len(s); i++ {
		a := s[i-1]
		if a > r {
			r = a
		}
	}
	return r
}