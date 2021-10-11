package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready."

func main() {

	//one way to declare variables - 2 steps
	var firstNumber int
	firstNumber = 2

	// another way, declare type and name and assign value
	var secondNumber = 5

	//one step variable:declare name, assign value, and let go figure out type
	subtraction := 7

	var answer = firstNumber*secondNumber - subtraction

	reader := bufio.NewReader((os.Stdin))
	fmt.Println("Guess the Number Game")
	fmt.Println("")

	fmt.Println("Guess a number between one and ten", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by ", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply that number by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the reusult by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	fmt.Println("Is the answer", answer, "? :)")
}
