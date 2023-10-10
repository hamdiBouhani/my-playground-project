package cmd

import (
	"fmt"
	"os"

	"github.com/hamdiBouhani/my-playground-project/config"
	"github.com/hamdiBouhani/my-playground-project/server"
	"github.com/spf13/cobra"
)

func NewRestCmd(logging *config.LoggerClient) *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "Run rest server",
		Long:  ``,
		Run: func(commandServe *cobra.Command, args []string) {
			err := server.NewHttpService(logging).Start()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}
}
