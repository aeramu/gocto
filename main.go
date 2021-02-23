package main

import (
	"github.com/spf13/cobra"
	"gocto/cmd/generate"
	initcmd "gocto/cmd/init"
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
