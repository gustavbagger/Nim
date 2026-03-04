package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

/*
	Add features:

- toggle functionality for 'You have lost already, I can force a win' message
- repeat same setup as last round
*/

type player struct {
	name string
	id   int
	wins int
}

func players() []player {
	fmt.Println("Who is playing: (one name for pve or two names for pvp)")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return players()
	}
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")
	if len(args) == 1 {
		return []player{
			{"computer", 0, 0},
			{args[0], 1, 0},
		}
	} else if len(args) == 2 {
		var playerSlice []player
		for i, name := range args {
			playerSlice = append(playerSlice, player{name, i, 0})
		}
		return playerSlice
	} else {
		fmt.Println("Currently we only support one- or two-player games")
		return players()
	}

}

func setup() []int {
	fmt.Println("Enter desired game setup: (d for default)")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return setup()
	}
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")

	var argInts []int
	if args[0] == "d" {
		return []int{7, 5, 3, 1}
	}
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

func (state *gamestate) playerMove(player player) {
	fmt.Printf("%v's turn: ", player.name)
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Usage: <row> <tiles removing>", err)
		state.playerMove(player)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")
	if len(args) != 2 {
		fmt.Println("Usage: <row> <tiles removing>")
		state.playerMove(player)
		return
	}
	row, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("invalid row")
		state.playerMove(player)
		return
	}
	removing, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("invalid tile removal")
		state.playerMove(player)
		return
	}
	if !checkValidMove(row-1, removing, state.columns) {
		fmt.Println("Invalid move")
		state.playerMove(player)
		return
	}

	state.move(row-1, removing)
	state.whoJustPlayed = player
	printGamestate(state.columns)
}

func (state *gamestate) displayWin() []player {
	fmt.Println("=======================================================================")
	fmt.Printf("%v wins!\n", state.whoJustPlayed.name)

	winningPlayerID := state.whoJustPlayed.id
	state.players[winningPlayerID].wins += 1

	fmt.Printf(
		"%s: %v wins --- %s: %v wins\n",
		state.players[0].name, state.players[0].wins,
		state.players[1].name, state.players[1].wins,
	)

	fmt.Println("=======================================================================")
	return state.players
}

func whoStarts(players []player) player {
	fmt.Println("Who starts? (computer/player name)")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("type one name", err)
		return whoStarts(players)
	}

	// remove the delimeter from the string
	nameOfStartingPlayer := strings.TrimSpace(input)
	for _, player := range players {
		if player.name == nameOfStartingPlayer {
			return player
		}
	}
	fmt.Println("could not find player")
	return whoStarts(players)
}

func playAgain() bool {
	fmt.Printf("Would you like to play again? (y/n) ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		return playAgain()
	}
	input = strings.TrimSpace(input)
	if input == "y" {
		return true
	} else if input == "n" {
		fmt.Println("\n=======================================================================")
		fmt.Println("See you next time")
		fmt.Println("=======================================================================")
		os.Exit(0)
		return false
	} else {
		return playAgain()
	}
}

func main() {
	fmt.Println("=======================================================================")
	fmt.Println(
		"--Welcome to Nim--\n",
		"This game consists of several rows of tiles and is turn based.\n",
		"Players alternate by removing tiles from a single row at a time.\n",
		"The winner is the player who removes the last tile from the entire board.",
	)
	fmt.Println("=======================================================================")
	playerMap := players()
	fmt.Println("=======================================================================")

	// Setup interrupt signal handler
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		fmt.Println("\n=======================================================================")
		fmt.Println("See you next time")
		fmt.Println("=======================================================================")
		os.Exit(0)
	}()

	for i := 1; ; i++ {
		emptyGamestate := gamestate{
			columns:        setup(),
			whoJustPlayed:  player{},
			players:        playerMap,
			startingPlayer: whoStarts(playerMap),
		}
		state := &emptyGamestate

		fmt.Println("=======================================================================")
		fmt.Printf("Round %v has been initialised in the following state:\n", i)
		printGamestate(state.columns)
		fmt.Println("=======================================================================")

	roundloop:
		for {
			if state.startingPlayer.name == "computer" {
				state.computerMove(state.startingPlayer)
			} else {
				state.playerMove(state.startingPlayer)
			}
			if checkWin(state.columns) {
				playerMap = state.displayWin()
				break roundloop
			}
			for _, player := range state.players {
				if player.name != state.startingPlayer.name {
					if player.name == "computer" {
						state.computerMove(player)
					} else {
						state.playerMove(player)
					}
					if checkWin(state.columns) {
						playerMap = state.displayWin()
						break roundloop
					}
				}
			}
		}
		playAgain()

	}
}
