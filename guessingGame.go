package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

// Function to play the game
func playGame() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("If your guess is within 5 of the number, you win!")

	// Generate a random number between 1 and 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1
	attempts := 0
	maxAttempts := 7 // Allow up to 7 attempts

	for attempts < maxAttempts {
		// Take user input
		fmt.Printf("Attempt %d/%d: Make a guess: ", attempts+1, maxAttempts)
		var input string
		fmt.Scanln(&input)

		// Convert input to integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		// Check the guess
		if guess == secretNumber {
			fmt.Printf("Congratulations! You guessed the exact number in %d attempts!\n", attempts+1)
			break
		} else if abs(guess-secretNumber) <= 5 {
			fmt.Printf("Great job! Your guess %d is within 5 of the correct number (%d). You win!\n", guess, secretNumber)
			break
		} else if guess < secretNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}

		attempts++
	}

	// If the user runs out of attempts
	if attempts == maxAttempts && abs(secretNumber-secretNumber) > 5 {
		fmt.Printf("Sorry, you've used all your attempts. The correct number was %d.\n", secretNumber)
	}

	// Ask to play again
	var playAgain string
	fmt.Print("Do you want to play again? (yes/no): ")
	fmt.Scanln(&playAgain)

	if playAgain == "yes" {
		playGame()
	} else {
		fmt.Println("Thanks for playing! Goodbye!")
	}
}

// Helper function to calculate the absolute difference
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Main function to start the game
func main() {
	// Start the game
	playGame()
}
