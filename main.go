package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

/* Add features:
- toggle functionality for 'You have lost already, I can force a win' message
- repeat same setup as last round
*/

func setup() []int {
	fmt.Println("Enter desired game setup: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return setup()
	}
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	var argInts []int
	for _, str := range args {
		argInt, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("invalid entry")
			return setup()
		}
		argInts = append(argInts, argInt)
	}
	return argInts

}

func main() {
	fmt.Println("=========================================================")
	fmt.Println("Welcome to Nim")
	fmt.Println("=========================================================")

	currentScore := []int{0, 0}

	var Winner int

	// Setup interrupt signal handler
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		fmt.Println("\n=========================================================")
		fmt.Println("See you next time")
		fmt.Println("=========================================================")
		os.Exit(0)
	}()
gameloop:
	for i := 1; ; i++ {

		emptyGamestate := gamestate{
			columns: setup(),
		}
		state := &emptyGamestate

		fmt.Println("=========================================================")
		fmt.Printf("Round %v has been initialised in the following state:\n", i)
		printGamestate(state.columns)
		fmt.Println("=========================================================")

	roundloop:
		for {
			if i%2 == 0 {
				state.computerMove()
				if checkWin(state.columns) {
					Winner = 0
					break roundloop
				}
				printGamestate(state.columns)
			}
		playerturn:
			for {
				fmt.Printf("Your move: ")
				reader := bufio.NewReader(os.Stdin)
				// ReadString will block until the delimiter is entered
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Usage: <row> <tiles removing>", err)
					continue playerturn
				}

				// remove the delimeter from the string
				input = strings.TrimSuffix(input, "\n")
				args := strings.Split(input, " ")
				if len(args) != 2 {
					fmt.Println("Usage: <row> <tiles removing>")
					continue playerturn
				}
				row, err := strconv.Atoi(args[0])
				if err != nil {
					fmt.Println("invalid row")
					continue playerturn
				}
				removing, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println("invalid tile removal")
					continue playerturn
				}
				if !checkValidMove(row-1, removing, state.columns) {
					fmt.Println("Invalid move")
					continue playerturn
				}

				state.move(row-1, removing)
				if checkWin(state.columns) {
					Winner = 1
					break roundloop
				}
				printGamestate(state.columns)
				break playerturn
			}
			if i%2 != 0 {
				state.computerMove()
				if checkWin(state.columns) {
					Winner = 0
					break roundloop
				}
				printGamestate(state.columns)
			}

		}

		fmt.Println("=========================================================")
		if Winner == 0 {
			fmt.Println("Computer wins :/")
			currentScore[0] += 1
		} else {
			fmt.Println("You win!!")
			currentScore[1] += 1
		}
		fmt.Printf("Computer: %v Player: %v\n", currentScore[0], currentScore[1])
		fmt.Println("=========================================================")
		for {
			fmt.Printf("Would you like to play again? (y/n) ")
			reader := bufio.NewReader(os.Stdin)
			// ReadString will block until the delimiter is entered
			input, err := reader.ReadString('\n')
			if err != nil {
				continue
			}

			if input == "y\n" {
				continue gameloop
			} else {
				fmt.Println("\n=========================================================")
				fmt.Println("See you next time")
				fmt.Println("=========================================================")
				os.Exit(0)
			}
		}
	}

}
