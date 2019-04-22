package engine

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitProject() {
	projectName := viper.GetString("init")
	projectName = fmt.Sprintf("./%s", projectName)
	mainFileName := fmt.Sprintf("%s/main.go", projectName)
	if err := os.Mkdir(projectName, os.ModePerm); err != nil {
		println(err.Error())
		return
	}
	file, err := os.OpenFile(mainFileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		println(err)
		return
	}
	file.Write([]byte("OK"))
	file.Close()
}
