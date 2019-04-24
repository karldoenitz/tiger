package main

import (
	"github.com/karldoenitz/tiger/cmd"
	"github.com/karldoenitz/tiger/engine"
)

func main() {
	cmd.InitCmd()
	engine.Drive()
}
