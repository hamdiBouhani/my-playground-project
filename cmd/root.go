package cmd

import "github.com/spf13/cobra"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my-playground-project",
	Short: "my playground project",
}

// registerCommands is where all new command should be added
func registerCommands() {
	rootCmd.AddCommand(NewRestCmd())
}

// Execute adds all child command to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	registerCommands()
	cobra.CheckErr(rootCmd.Execute())
}
