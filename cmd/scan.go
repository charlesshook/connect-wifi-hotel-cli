package cmd

import (
	"fmt"
	"strings"

	"github.com/charlesshook/connect-wifi-hotel-cli/internal/utils"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans for advaiable networks.",
	Long:  "Will scan for the advailable neworks and display their SSID, Signal strength and security protocol.",
	Run:   scan,
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

func scan(cmd *cobra.Command, args []string) {
	fmt.Println("Scanning for available WiFi networks...")

	output, err := utils.RunCommand("nmcli -t -f ssid,signal,security dev wifi list")

	if err != nil {
		fmt.Printf("failed to scan WiFi networks: %v", err)
	}

	if output == "" {
		fmt.Println("No WiFi networks found.")
	}

	fmt.Println("Available WiFi Networks:")
	fmt.Println("SSID|Signal|Security")
	fmt.Println("----------------------------------------------------------------------")
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")

		if len(fields) >= 3 {
			ssid := fields[0]
			if ssid == "" {
				ssid = "<hidden>"
			}

			signal := fields[1]
			security := fields[2]

			if security == "" {
				security = "<not listed>"
			}

			fmt.Printf("%s\t\t%s\t%s\n", ssid, signal, security)
		}
	}
}
