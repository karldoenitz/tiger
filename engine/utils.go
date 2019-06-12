package engine

import (
	"bytes"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		println(err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func GetGoPath() string {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	return goPath
}

func ShellExecuteEngine(command string) (error, string, string) {
	ShellToUse := "bash"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if runtime.GOOS == "windows" {
		ShellToUse = "cmd"
	}
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
