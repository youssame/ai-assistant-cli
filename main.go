package main

import (
	"github.com/spf13/cobra"
	"github.com/youssame/assistant-cli/cmd/ai"
	"github.com/youssame/assistant-cli/cmd/vpn"
	"log"
	"os"
)

var RootCmd = &cobra.Command{
	Use:     "foo",
	Version: "0.1.0",
	Short:   "My personal CLI assistant",
}

func main() {
	checkEnvVariables()
	RootCmd.AddCommand(vpn.Cmd, ai.Cmd)
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func checkEnvVariables() {
	appName := os.Getenv("APP_NAME")
	vpnHost := os.Getenv("VPN_HOST")
	ciscoDir := os.Getenv("CISCO_BIN_DIR")
	llmHost := os.Getenv("LLM_HOST")
	llmModel := os.Getenv("LLM_MODEL")

	if appName == "" || vpnHost == "" || ciscoDir == "" || llmHost == "" || llmModel == "" {
		log.Fatal("Error while loading the env variables")
	}
}
