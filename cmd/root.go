package cmd

import (
	"log"

	"github.com/hamdiBouhani/my-playground-project/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my-playground-project",
	Short: "my playground project",
}

// registerCommands is where all new command should be added
func registerCommands(logging *config.LoggerClient) {
	rootCmd.AddCommand(NewRestCmd(logging))
	rootCmd.AddCommand(NewMigrateCmd())
}

// Execute adds all child command to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	logging, err := config.NewLoggerClient()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	registerCommands(logging)
	cobra.CheckErr(rootCmd.Execute())
}
