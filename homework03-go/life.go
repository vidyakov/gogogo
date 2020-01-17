package main

import (
	"math/rand"
	"fmt"
	"time"
)


var rows int
var cols int
var generations int
var randomize bool = true
//var currGeneration ?
//var maxgeneration ?

func createGrid(randomize bool) [][]int {
	var matrix [][]int
	for index, value := range

	if !randomize {
		return matrix
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello!")
}
