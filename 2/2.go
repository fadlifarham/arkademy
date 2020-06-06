package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Give me number : ")
	fmt.Scanf("%d", &n)

	isEven := n%2 == 0

	switch isEven {
	case true:
		fmt.Println("The input must be an odd number")
	case false:
		printImage(n)
	default:
		fmt.Println("Error")
	}
}

func printImage(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 || j == n-1 || i == n/2 {
				fmt.Print("* ")
			} else {
				fmt.Print("= ")
			}
		}
		fmt.Println("")
	}
}
