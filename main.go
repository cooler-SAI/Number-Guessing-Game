package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Create a random source using the current time for randomness
	randSource := rand.NewSource(time.Now().UnixNano())
	// Create a new random number generator using the random source
	randGenerator := rand.New(randSource)
	// Create a buffered reader to read input from the user
	reader := bufio.NewReader(os.Stdin)

	// Print welcome message and rules of the game
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 20.")
	fmt.Println("You need to guess the number correctly within the allowed chances.")

	for {
		// Display difficulty options to the user
		fmt.Println("\nPlease select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")
		fmt.Print("Enter your choice: ")
		// Read and trim the user's choice
		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		choice = strings.TrimSpace(choice)

		var chances int
		// Assign the number of chances based on the user's choice
		switch choice {
		case "1":
			chances = 10
		case "2":
			chances = 5
		case "3":
			chances = 3
		default:
			// Handle invalid input and restart the loop
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		// Generate a random number between 1 and 20 using the new random generator
		number := randGenerator.Intn(20) + 1
		fmt.Printf("\nGreat! You have selected difficulty with %d chances.\n", chances)
		fmt.Println("Let's start the game!")

		// Loop for the number of allowed attempts
		for attempts := 1; attempts <= chances; attempts++ {
			// Prompt the user for their guess
			fmt.Printf("Enter your guess (%d remaining): ", chances-attempts+1)
			guessStr, _ := reader.ReadString('\n')
			guessStr = strings.TrimSpace(guessStr)
			// Convert the user's input to an integer
			guess, err := strconv.Atoi(guessStr)
			if err != nil || guess < 1 || guess > 100 {
				// Handle invalid input
				fmt.Println("Please enter a valid number between 1 and 100.")
				continue
			}

			if guess == number {
				// Correct guess
				fmt.Printf("Congratulations! You guessed the correct number %d in %d attempts.\n", number, attempts)
				break
			} else if guess < number {
				// Guess is too low
				fmt.Println("Incorrect! The number is greater.")
			} else {
				// Guess is too high
				fmt.Println("Incorrect! The number is less.")
			}

			// If the user runs out of chances, reveal the number
			if attempts == chances {
				fmt.Printf("You've run out of chances! The correct number was %d.\n", number)
			}
		}

		// Ask the user if they want to play again
		fmt.Print("Do you want to play again? (yes/no): ")
		response, _ := reader.ReadString('\n')
		response = strings.ToLower(strings.TrimSpace(response))
		if response != "yes" {
			// Exit the game loop if the user does not want to play again
			fmt.Println("Thanks for playing! Goodbye!")
			break
		}
	}
}
