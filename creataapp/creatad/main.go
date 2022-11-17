package main

import (
	"os"

	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/creataapp/creatad/cmd"
	"github.com/creatachain/creata-sdk/server"
	svrcmd "github.com/creatachain/creata-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, creataapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
