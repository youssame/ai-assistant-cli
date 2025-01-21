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
	hostAddress := os.Getenv("ASSISTANT_VPN_HOST")
	ciscoBin := os.Getenv("ASSISTANT_CISCO_BIN_DIR")
	cmd := exec.Command(ciscoBin, "connect", "-s", "connect", hostAddress)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error executing the VPN connect command.")
	}
}

func disconnect() {
	ciscoBin := os.Getenv("ASSISTANT_CISCO_BIN_DIR")
	cmd := exec.Command(ciscoBin, "disconnect")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error executing the VPN disconnect command.")
	}
}

func state() {
	ciscoBin := os.Getenv("ASSISTANT_CISCO_BIN_DIR")
	cmd := exec.Command(ciscoBin, "state")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error executing the VPN state command.")
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
	vpnStateCmd := &cobra.Command{
		Use:     "s",
		Version: "0.1.0",
		Short:   "state of the Cisco secure vpn",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Loading the VPN state....")
			state()
		},
	}
	Cmd.AddCommand(vpnConnectCmd, vpnDisconnectCmd, vpnStateCmd)
}
