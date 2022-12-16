package main

import (
	"fmt"

	"github.com/litusluca/litusluca.github.io/src/app"
)

func main() {
	fmt.Println("Hello Wordl!")
	game := app.App("game", new(MyApp))

	game.Run()
}