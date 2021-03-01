package main

import (
	"github.com/aeramu/gocto/cmd/generate"
	initcmd "github.com/aeramu/gocto/cmd/init"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use: "gocto",
	}
	generateCmd = &cobra.Command{
		Use: "generate",
		Run: generate.Run,
	}
	initCmd = &cobra.Command{
		Use: "init",
		Run: initcmd.Run,
	}
)

func init() {
	cmd.AddCommand(generateCmd, initCmd)
}

func main() {
	cmd.Execute()
}
