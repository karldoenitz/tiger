package engine

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var projectionName string

// InitProject 初始化项目
func InitProject() {
	fmt.Println("init projection start")
	projectName := viper.GetString("init")
	projectionName = projectName
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
	mainCode := fmt.Sprintf(main, ".")
	currentPath := GetCurrentDirectory()
	goPath := GetGoPath()
	if goPath+"/src" == currentPath {
		mainCode = fmt.Sprintf(main, viper.GetString("init"))
	}
	_, err = file.Write([]byte(mainCode))
	if err != nil {
		println(err.Error())
	}
	err = file.Close()
	if err != nil {
		println(err.Error())
	}
	createConfigFile()
	createHandlers()
	createTemplates()
	fmt.Printf("projection %s created\n", viper.GetString("init"))
}

// createConfigFile 创建配置文件
func createConfigFile() {
	projectName := viper.GetString("init")
	configName := fmt.Sprintf("./%s/server.yaml", projectName)
	file, err := os.OpenFile(configName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		println(err)
		return
	}
	projectPath := GetCurrentDirectory() + "/" + projectionName
	config := fmt.Sprintf(configuration, projectPath)
	_, err = file.Write([]byte(config))
	if err != nil {
		println(err.Error())
	}
	err = file.Close()
	if err != nil {
		println(err.Error())
	}
}

// createHandlers 创建handlers文件及文件夹
func createHandlers() {
	projectName := viper.GetString("init")
	handlers := fmt.Sprintf("./%s/handlers", projectName)
	if err := os.Mkdir(handlers, os.ModePerm); err != nil {
		println(err.Error())
		return
	}
	handlerName := fmt.Sprintf("%s/pinghandler.go", handlers)
	file, err := os.OpenFile(handlerName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		println(err)
		return
	}
	_, err = file.Write([]byte(fmt.Sprintf(handler, projectName)))
	if err != nil {
		println(err.Error())
	}
	err = file.Close()
	if err != nil {
		println(err.Error())
	}
}

// createTemplates 创建模板
func createTemplates() {
	staticPath := GetCurrentDirectory() + "/" + projectionName + "/static"
	if err := os.Mkdir(staticPath, os.ModePerm); err != nil {
		println(err.Error())
		return
	}
	projectPath := GetCurrentDirectory() + "/" + projectionName + "/static/front.html"
	file, err := os.OpenFile(projectPath, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		println(err)
		return
	}
	_, err = file.Write([]byte(template))
	if err != nil {
		println(err.Error())
	}
	err = file.Close()
	if err != nil {
		println(err.Error())
	}
}
