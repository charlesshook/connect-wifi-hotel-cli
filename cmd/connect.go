package cmd

import (
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Will connect to a network.",
	Long:  "Will connect to a network. Can handle captive portals. Can take in a name and room number. Will also handle a network with a password. If auto captive portal does not work then it will allow for manualy intervection.",
	Run:   connect,
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

func connect(cmd *cobra.Command, args []string) {

}
