package main

import (
	"os"

	"github.com/hirokryptor/xyzchain/cmd/xyzchaind/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
