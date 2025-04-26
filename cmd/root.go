package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "connect-hotel-wifi",
	Short: "Helpful tool to connect to hotel wifi networks.",
	Long: `This tool helps you connect to hotel wifi networks.
	It can help deal with those pesky captive portals.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
