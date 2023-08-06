package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hmdshariati/go-practice/cmd"
	"github.com/spf13/cobra"
)

// import "fmt"

func main() {
	const description = "test command"
	
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root.AddCommand(
		cmd.Server{}.Command(trap),

	)

}