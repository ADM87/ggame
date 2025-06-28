package main

import (
	"github.com/ADM87/ggame/src/game"
)

var version = "0.0.0-unreleased"

func main() {
	if err := game.Start(version); err != nil {
		panic(err)
	}
}
