package main

import (
	"github.com/aeramu/gocto/cmd/generate"
	initcmd "github.com/aeramu/gocto/cmd/init"
	"github.com/aeramu/gocto/cmd/service"
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
	serviceCmd = &cobra.Command{
		Use: "service",
	}
	serviceAddCmd = &cobra.Command{
		Use: "add",
		Run: service.Add,
	}
)

func init() {
	serviceCmd.AddCommand(serviceAddCmd)
	cmd.AddCommand(generateCmd, initCmd, serviceCmd)
}

func main() {
	cmd.Execute()
}
