package main

import "fmt"

type Doctor struct {
	number    int
	name      string
	Companion []string
}

func main() {
	a := Doctor{
		number: 3,
		name:   "jon pee",
		Companion: []string{
			"liz",
			"jo",
			"sarah",
		},
	}
	fmt.Println(a)
	fmt.Println("number -", a.number)
}
