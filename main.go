package main

import (
	"github.com/spf13/cobra"
	"github.com/youssame/assistant-cli/cmd/ai"
	"github.com/youssame/assistant-cli/cmd/team"
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
	appName := os.Getenv("ASSISTANT_APP_NAME")
	db := os.Getenv("ASSISTANT_DB_HOST")
	if vpnHost == "" {
		log.Println("Error: The environment variables are not set.")
		log.Println("export ASSISTANT_DB_HOST=<value>")
	}
	if appName == "" {
		log.Println("Error: The environment variables are not set.")
		log.Println("export ASSISTANT_APP_NAME=<value>")
	}
	if ciscoDir == "" {
		log.Println("Error: The environment variables are not set.")
		log.Println("export ASSISTANT_CISCO_BIN_DIR=<value>")
	}
	if llmHost == "" {
		log.Println("Error: The environment variables are not set.")
		log.Println("export ASSISTANT_LLM_HOST=<value>")
	}
	if llmModel == "" {
		log.Println("Error: The environment variables are not set.")
		log.Println("export ASSISTANT_LLM_MODEL=<value>")
	}
	if db == "" {
		log.Println("Error: The environment variables are not set.")
		log.Println("export ASSISTANT_DB_HOST=<value>")
	}
	if vpnHost == "" || ciscoDir == "" || llmHost == "" || llmModel == "" || db == "" || appName == "" {
		os.Exit(1)
	}
}
