package engine

import "github.com/spf13/viper"

func Drive() {
	projectName := viper.GetString("init")
	version := viper.GetBool("version")
	run := viper.GetBool("run")
	update := viper.GetString("update")
	if projectName != "" {
		InitProject()
	}
	if version {
		showVersionInfo()
	}
	if run {
		if err, out, errOut := ShellExecuteEngine("go run"); err != nil {
			println(errOut)
		} else {
			println(out)
		}
	}
	if update != "" {
		updateFramework()
	}
}
