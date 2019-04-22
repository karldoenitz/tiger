package main

import (
	"fmt"
	"github.com/spf13/viper"
	"tiger/cmd"
)

func main() {
	cmd.InitCmd()
	fmt.Println("Tiger OK")
	fmt.Println(viper.GetString("mode"))
}
