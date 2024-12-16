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

func determineChances(choice string) (int, error) {
	switch choice {
	case "1":
		return 10, nil
	case "2":
		return 5, nil
	case "3":
		return 3, nil
	default:
		return 0, fmt.Errorf("invalid choice")
	}
}

func checkGuess(guess, number int) string {
	if guess == number {
		return "correct"
	} else if guess < number {
		return "greater"
	} else {
		return "less"
	}
}

func main() {
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You need to guess the number correctly within the allowed chances.")

	for {
		fmt.Println("\nPlease select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		chances, err := determineChances(choice)
		if err != nil {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		number := randGenerator.Intn(100) + 1
		fmt.Printf("\nGreat! You have selected difficulty with %d chances.\n", chances)
		fmt.Println("Let's start the game!")

		for attempts := 1; attempts <= chances; attempts++ {
			fmt.Printf("Enter your guess (%d remaining): ", chances-attempts+1)
			guessStr, _ := reader.ReadString('\n')
			guessStr = strings.TrimSpace(guessStr)

			guess, err := strconv.Atoi(guessStr)
			if err != nil || guess < 1 || guess > 100 {
				fmt.Println("Please enter a valid number between 1 and 100.")
				continue
			}

			result := checkGuess(guess, number)
			if result == "correct" {
				fmt.Printf("Congratulations! You guessed the correct number %d in %d attempts.\n",
					number, attempts)
				break
			} else if result == "greater" {
				fmt.Println("Incorrect! The number is greater.")
			} else {
				fmt.Println("Incorrect! The number is less.")
			}

			if attempts == chances {
				fmt.Printf("You've run out of chances! The correct number was %d.\n", number)
			}
		}

		fmt.Print("Do you want to play again? (yes/no): ")
		response, _ := reader.ReadString('\n')
		response = strings.ToLower(strings.TrimSpace(response))
		if response != "yes" {
			fmt.Println("Thanks for playing! Goodbye!")
			break
		}
	}
}
