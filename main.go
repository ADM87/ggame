package main

import (
	ggame "github.com/ADM87/ggame/src"
)

var version = "0.0.0-unreleased"

func main() {
	if err := ggame.Boot(version); err != nil {
		panic(err)
	}
}
