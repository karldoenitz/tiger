package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func init() {
	pflag.StringP("init", "i", "", "初始化一个Tigo项目 / init a Tigo project")
	pflag.StringP("update", "u", "", "更新tiger或者Tigo / update tiger or Tigo")
	pflag.BoolP("version", "v", false, "是否显示版本号 / show version")
	pflag.BoolP("run", "r", false, "运行当前项目 / run projection")
	pflag.Usage = usage
	pflag.ErrHelp = errors.New("")
	pflag.Parse()
	if e := viper.BindPFlags(pflag.CommandLine); e != nil {
		println(e.Error())
	}
	initConfig()
}

func usage() {
	_, err := fmt.Fprintf(os.Stderr,
		`
Tigo Commandline Tool
	
Usage: tiger [Options]

Options:
	-i, --init projection_name 初始化一个Tigo项目 / init a Tigo projection
	-u, --update tiger/tigo    更新tiger工具或者Tigo框架 / update tiger or Tigo
	-v, --version              查看Tigo以及tiger的版本号 / show version of Tigo or tiger
	-r, --run                  在当前项目路径下运行项目 / run current projection
	-h, --help                 查看帮助文档 / show help
    `)
	if err != nil {
		println(err.Error())
	}
}

func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}

func InitCmd() {

}
