package main

import (
	"fmt"
	"runtime"

	"github.com/thinkgos/sharp/builder"
)

func main() {
	fmt.Println("App Name: " + builder.Name)
	fmt.Println("Build Time: " + builder.BuildTime)
	fmt.Println("Git Commit: " + builder.GitCommit)
	fmt.Println("Git Full Commit: " + builder.GitFullCommit)
	fmt.Println("Version: " + builder.Version)
	fmt.Println("API Version: " + builder.APIVersion)
	fmt.Println("Model: " + builder.Model)

	fmt.Println("Go version: " + runtime.Version())
	fmt.Println("Os/Arch: " + runtime.GOOS + "/" + runtime.GOARCH)
}
