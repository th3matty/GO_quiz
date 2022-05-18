package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to my quiz game!")
	fmt.Printf("Please enter your Name: \n") // breaks to new line "\n"
	var name string
	fmt.Scan(&name) // Scan(&) -> stores a value in ram

	fmt.Printf("Hello, %v, welcome to the game!", name) // Printf, otherwise we cannot use %v
	fmt.Printf("Please enter your age: \n")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("roger that, lets go")
	} else {
		fmt.Println("sorry, you`re to young")
		return
	}

	score := 0
	num_questions := 2

	fmt.Printf("What is better? RTX 3080 or RTX 3090? \n")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	res1 := strings.ToLower(answer)
	res2 := strings.ToLower(answer2)

	if res1+" "+res2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? \n")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Inccorect")
	}

	fmt.Printf("You scored %v out of %v \n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100

	fmt.Printf("You scored: %v%%.", percent)
}
