# Golang
This is a repo for my Golang projects.
package main

import "fmt"

const (
	rainy  = 1
	summer = 3
	yearly = 4
)

func main() {
	books := [...]string{
		"Kaanthi",
		"wings of fire",
		"life",
		"Kanthi Second edition",
	}
	var (
		rbooks [rainy]string
		sbooks [summer]string
	)
	rbooks[0] = books[0]
	for i := range sbooks {
		sbooks[i] = books[i+1]
	}
	fmt.Printf("Rainy books  : %#q\n", rbooks)
	fmt.Printf("Summer books : %#q\n", sbooks)

}
