package engine

import (
	"encoding/json"
	"fmt"
	"github.com/karldoenitz/Tigo/TigoWeb"
	"github.com/karldoenitz/Tigo/request"
	"github.com/spf13/viper"
)

type VersionInfo struct {
	Author string    `json:"author"`
	Tigo   FrameWork `json:"tigo"`
	Tiger  FrameWork `json:"tiger"`
}

type FrameWork struct {
	Version string `json:"version"`
	Link    string `json:"link"`
	Website string `json:"website"`
}

// showVersionInfo 显示版本信息
func showVersionInfo() {
	fmt.Printf("Tigo  local  version: %s\n", TigoWeb.Version)
	fmt.Printf("tiger local  version: %s\n", Version)
	versionInfo := VersionInfo{}
	response, err := request.Get("https://karldoenitz.github.io/Tigo-EN/static/data/info.json")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Tigo  latest version: unknown")
		fmt.Println("tiger latest version: unknown")
		return
	}
	err = json.Unmarshal(response.Content, &versionInfo)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Tigo  latest version: %s\n", versionInfo.Tigo.Version)
	fmt.Printf("tiger latest version: %s\n", versionInfo.Tiger.Version)
}

func updateFramework() {
	name := viper.GetString("init")
	shell := ""
	switch name {
	case "tiger":
		shell = "go get -u github.com/karldoenitz/tiger"
		break
	case "tigo":
		shell = "go get -u github.com/karldoenitz/Tigo/..."
		break
	}
	if err, out, errOut := ShellExecuteEngine(shell + "tiger"); err != nil {
		println(errOut)
	} else {
		println(out)
	}
}
