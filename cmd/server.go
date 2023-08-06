package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.run(trap)
	}

	return &cobra.Command{
		Use: "server",

		Short: "run frist command",

		Run: run,
	}
}

func (cmd Server) run(trap chan os.Signal) {
	fmt.Printf("Hello World")
}
