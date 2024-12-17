package engine

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func PrintWelcomeMessage() {
	fmt.Println("=======================================")
	fmt.Println("||      Welcome to the Game Hub!     ||")
	fmt.Println("||                                   ||")
	fmt.Println("||   Guess the Number between 1-100  ||")
	fmt.Println("||   Choose a difficulty level and   ||")
	fmt.Println("||   try your best to win the game!  ||")
	fmt.Println("=======================================")
}

func TestHuman() {
	fmt.Println("Human test")
}

func HandleExitSignal() {

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exitSignal
		fmt.Println("\nExit signal received")
		os.Exit(0)
	}()
}
