package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func init() {
	pflag.StringP("init", "i", "greatProject", "init a Tigo project / 初始化一个Tigo项目")
	pflag.Bool("version", false, "show version / 是否显示版本号")
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
	-i, --inti projection_name init a Tigo projection / 初始化一个Tigo项目
	-v, --version              show version of Tigo or tiger / 查看Tigo以及tiger的版本号
	-h, --help                 show help / 查看帮助文档
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
