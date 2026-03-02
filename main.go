package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Add features:
- toggle functionality for 'You have lost already, I can force a win' message
- choose different starting setups
*/

func main() {
	fmt.Println("=========================================================")
	fmt.Println("Welcome to Nim")
	fmt.Println("=========================================================")

	var state *gamestate

gameSetup:
	for {
		fmt.Println("Enter desired game setup: ")
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		args := strings.Split(input, " ")

		var argInts []int
		for _, str := range args {
			argInt, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("invalid entry")
				continue gameSetup
			}
			argInts = append(argInts, argInt)
		}
		state = gamestateNew(argInts...)

		break gameSetup

	}
	fmt.Println("=========================================================")
	fmt.Println("Game has been initialised in the following state:")
	printGamestate(state.columns)
	fmt.Println("=========================================================")

	//gameloop
	for {
		for {
			fmt.Printf("Your move: ")
			reader := bufio.NewReader(os.Stdin)
			// ReadString will block until the delimiter is entered
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
				continue
			}

			// remove the delimeter from the string
			input = strings.TrimSuffix(input, "\n")
			args := strings.Split(input, " ")
			if len(args) != 2 {
				fmt.Println("Usage: <row> <tiles removing>")
				continue
			}
			row, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("invalid row")
				continue
			}
			removing, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("invalid tile removal")
				continue
			}
			state.move(row-1, removing)
			printGamestate(state.columns)

			state.computerMove()
			printGamestate(state.columns)
			break

		}

	}
	/*
			//Exit programme proceedure with ctrl+C
			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, os.Interrupt)
			<-signalChan
			break
			fmt.Println("")
		fmt.Println("=========================================================")
			fmt.Println("See you next time")
		fmt.Println("=========================================================")
	*/
}
