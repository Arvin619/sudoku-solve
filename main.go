package main

import (
	"fmt"
	"os"

	"github.com/Arvin619/sudoku-solve/sudoku"
)

func main() {
	s := sudoku.New()

	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s.SetTableWithReader(file)
	s.Solve()
	fmt.Println(s)
}
