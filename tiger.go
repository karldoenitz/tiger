package main

import (
	"tiger/cmd"
	"tiger/engine"
)

func main() {
	cmd.InitCmd()
	engine.Drive()
}
