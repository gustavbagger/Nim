package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type gamestate struct {
	columns []int
}

func gamestateNew(columns ...int) *gamestate {
	return &gamestate{
		columns: columns,
	}
}

func sumWithoutCarry(columns []int) int {
	var tot int
	for _, column := range columns {
		tot = tot ^ column
	}
	return tot
}
func numNonZero(columns []int) int {
	nonEmpties := 0
	for _, column := range columns {
		if column != 0 {
			nonEmpties += 1
		}
	}
	return nonEmpties
}

func playRandomNonZero(columns []int) int {
	var nonEmptyIndexes []int
	for i, column := range columns {
		if column != 0 {
			nonEmptyIndexes = append(nonEmptyIndexes, i)
		}
	}
	randIndex := rand.Intn(len(nonEmptyIndexes))
	return nonEmptyIndexes[randIndex]

}

func findOptimalPlay(columns []int) (column int, removing int, err error) {
	sumOfColums := sumWithoutCarry(columns)

	// Computer is losing currently, choose random minimal move
	if sumOfColums == 0 {
		return playRandomNonZero(columns), 1, nil
	}
	// Computer is winning
	for i := range columns {
		column := columns[i]

		sumWithoutColumn := column ^ sumOfColums

		for x := 1; x <= column; x++ {
			if (column-x)^sumWithoutColumn == 0 {
				return i, x, nil
			}
		}
	}
	return 0, 0, errors.New("couldnt find optimal move")
}

func (state *gamestate) computerMove() error {
	column, removing, err := findOptimalPlay(state.columns)
	if err != nil {
		fmt.Println("Logic failure: ", err)
		return err
	}
	state.move(column, removing)
	fmt.Printf("Computer move: %v %v\n", column+1, removing)
	return nil
}

func (state *gamestate) move(column, removing int) {
	state.columns[column] -= removing
}

func checkValidMove(column, removing int, columns []int) bool {
	if len(columns) <= column {
		return false
	}
	return columns[column] >= removing
}

func checkWin(columns []int) bool {
	return numNonZero(columns) == 0
}
