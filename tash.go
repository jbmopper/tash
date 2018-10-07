package tash

import ()

const m = 62 // 26 lower case, 26 upper case, 10 digits

//Tash takes an integer and returns a string of lower and uppercase letters
//and numbers.  The goal is to compress an integer into a safe ASCII string
func Tash(i int) string {
	out := ""
	parse := func(ii int) string {
		s := ""
		switch {
		case ii < 0 || ii > 61:
			panic("Trying to parse an invalid integer...")
		case ii < 10:
			ii = ii + 48
		case ii < 37:
			ii = ii - 10 + 97
		case ii >= 37:
			ii = ii - 37 + 65
		default:
			panic("Math is hard")
		}
		s = string(ii)
		return s
	}
	if i == 0 {
		out = "0"
	}
	for i > 0 {
		q := i % m
		out = parse(q) + out
		i = i / m
	}
	return out
}
