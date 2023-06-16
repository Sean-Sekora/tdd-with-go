package main

import (
	"fmt"
	"tdd-with-go/domain"
)

func main() {
	word := domain.Word("cat")
	fmt.Println(word.Guess("tac"))
}
