package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

const m = "x = 42, y = 21\n"

func Printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	Printstr(m)
}
