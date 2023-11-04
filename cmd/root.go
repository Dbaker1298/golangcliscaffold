/*
Copyright © 2023 David Baker  <dpbaker1298@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose bool
var Debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "getsize",
	Short: "List the size of a local diretory.",
	Long: `This command will list the size of a local directory with several different options.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debug output in console output. (default: false)")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}


