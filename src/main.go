package main

import "fmt"

func main() {
	var score int
	fmt.Print("Enter your score: ")
	fmt.Scanln(&score)

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}
}
