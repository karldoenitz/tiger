package engine

import "github.com/spf13/viper"

func Drive() {
	projectName := viper.GetString("init")
	version := viper.GetBool("version")
	if projectName != "" {
		InitProject()
	}
	if version {
		showVersionInfo()
	}
}
