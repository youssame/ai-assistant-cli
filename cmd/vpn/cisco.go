package vpn

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var Cmd = &cobra.Command{
	Use:     "vpn",
	Version: "0.1.0",
	Short:   "Manage the Cisco secure vpn",
}

func connect() {
	hostAddress := os.Getenv("VPN_HOST")
	ciscoBin := os.Getenv("CISCO_BIN_DIR")
	cmd := exec.Command(ciscoBin, "connect", "-s", "connect", hostAddress)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error executing the VPN connect command.")
	}
}

func disconnect() {
	ciscoBin := os.Getenv("CISCO_BIN_DIR")
	cmd := exec.Command(ciscoBin, "disconnect")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error executing the VPN disconnect command.")
	}
}

func stats() {
	ciscoBin := os.Getenv("CISCO_BIN_DIR")
	cmd := exec.Command(ciscoBin, "stats")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error executing the VPN stats command.")
	}
}

func init() {
	vpnConnectCmd := &cobra.Command{
		Use:     "c",
		Version: "0.1.0",
		Short:   "connect the Cisco secure vpn",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Start connecting to VPN....")
			connect()
		},
	}
	vpnDisconnectCmd := &cobra.Command{
		Use:     "d",
		Version: "0.1.0",
		Short:   "disconnect the Cisco secure vpn",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Start disconnecting from VPN....")
			disconnect()
		},
	}
	vpnStatsCmd := &cobra.Command{
		Use:     "s",
		Version: "0.1.0",
		Short:   "stats of the Cisco secure vpn",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Loading the VPN stats....")
			stats()
		},
	}
	Cmd.AddCommand(vpnConnectCmd, vpnDisconnectCmd, vpnStatsCmd)
}
