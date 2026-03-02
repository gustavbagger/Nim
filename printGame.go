package main

import "fmt"

// Assuming gameloop has reaped empty colums
func printGamestate(columns []int) {

	for i, column := range columns {
		columnString := fmt.Sprintf("%v) ", i+1)
		for j := 0; j < column; j++ {
			columnString += "* "
		}
		fmt.Println(columnString)
	}
}
