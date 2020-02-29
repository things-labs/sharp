package main

import (
	"fmt"

	"github.com/thinkgos/assist/builder"
)

func main() {
	fmt.Println("Build Time: " + builder.BuildTime)
	fmt.Println("Git Commit: " + builder.GitCommit)
	fmt.Println("Git Full Commit: " + builder.GitFullCommit)
	fmt.Println("Version: " + builder.Version)
	fmt.Println("API Version: " + builder.APIVersion)
	fmt.Println("Model: " + builder.Model)
}
