package engine

import (
	"fmt"
	"testing"
)

func TestShowVersionInfo(t *testing.T) {
	showVersionInfo()
}

func TestGetCurrentDirectory(t *testing.T) {
	fmt.Printf("current path: %s", GetCurrentDirectory())
}
