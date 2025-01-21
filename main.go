package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/youssame/assistant-cli/cmd/ai"
	"github.com/youssame/assistant-cli/cmd/team"
	"github.com/youssame/assistant-cli/cmd/vpn"
	"os"
)

var RootCmd = &cobra.Command{
	Use:     "foo",
	Version: "0.1.0",
	Short:   "My personal CLI assistant",
}

func main() {
	checkEnvVariables()
	RootCmd.AddCommand(vpn.Cmd, ai.Cmd, team.Cmd)
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func checkEnvVariables() {
	vpnHost := os.Getenv("ASSISTANT_VPN_HOST")
	ciscoDir := os.Getenv("ASSISTANT_CISCO_BIN_DIR")
	llmHost := os.Getenv("ASSISTANT_LLM_HOST")
	llmModel := os.Getenv("ASSISTANT_LLM_MODEL")
	elastic := os.Getenv("ASSISTANT_DB_HOST")

	if vpnHost == "" || ciscoDir == "" || llmHost == "" || llmModel == "" || elastic == "" {
		fmt.Println("Error: The environment variables are not set.")
		fmt.Println("export ASSISTANT_DB_HOST=<value>")
		fmt.Println("export ASSISTANT_APP_NAME=<value>")
		fmt.Println("export ASSISTANT_VPN_HOST=<value>")
		fmt.Println("export ASSISTANT_CISCO_BIN_DIR=<value>")
		fmt.Println("export ASSISTANT_LLM_HOST=<value>")
		fmt.Println("export ASSISTANT_LLM_MODEL=<value>")
		os.Exit(1)
	}
}
