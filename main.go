package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/youssame/assistant-cli/cmd/ai"
	"github.com/youssame/assistant-cli/cmd/vpn"
	"log"
	"os"
)

var RootCmd = &cobra.Command{
	Use:     "assistant-cli",
	Version: "0.1.0",
	Short:   "My personal CLI assistant",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	RootCmd.AddCommand(vpn.Cmd, ai.Cmd)
	err = RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
