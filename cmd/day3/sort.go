package main

func sort(s [3]int) [3]int {
	if s[0] > s[2] {
		s[0], s[2] = s[2], s[0]
	}
	if s[0] > s[1] {
		s[0], s[1] = s[1], s[0]
	}
	if s[1] > s[2] {
		s[1], s[2] = s[2], s[1]
	}
	return s
}
