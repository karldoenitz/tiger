package engine

import (
	"Tigo/TigoWeb"
	"encoding/json"
	"fmt"
	"github.com/karldoenitz/Tigo/request"
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

func showVersionInfo() {
	versionInfo := VersionInfo{}
	response, err := request.Get("https://karldoenitz.github.io/Tigo-EN/static/data/info.json")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Tigo  local  version: %s\n", TigoWeb.Version)
		fmt.Println("Tigo  latest version: unknown")
		fmt.Printf("tiger local  version %s\n", Version)
		fmt.Println("tiger latest version: unknown")
		return
	}
	err = json.Unmarshal(response.Content, &versionInfo)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Tigo  local  version: %s\n", TigoWeb.Version)
	fmt.Printf("Tigo  latest version: %s\n", versionInfo.Tigo.Version)
	fmt.Printf("tiger local  version: %s\n", Version)
	fmt.Printf("tiger latest version: %s\n", versionInfo.Tiger.Version)
}
