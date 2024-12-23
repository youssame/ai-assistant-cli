package ai

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/youssame/assistant-cli/internal"
	"log"
)

var Cmd = &cobra.Command{
	Use:     "llm",
	Version: "0.1.0",
	Short:   "Manage requests to local LLM",
}

// reformulate a given text
func reformulate(message string) {
	res, err := internal.GenerateResponse(`reformulate this "` + message + `"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	fmt.Println("The result has been copied to the clipboard.")
}

// correct a given text
func correct(message string) {
	res, err := internal.GenerateResponse(`correct this "` + message + `"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	fmt.Println("The result has been copied to the clipboard.")
}

// ask a given question
func ask(message string) {
	res, err := internal.GenerateResponse(`answer this question "` + message + `?"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	fmt.Println("The result has been copied to the clipboard.")
}

// generate a given request (announcement, email, message...)
func generate(message string) {
	res, err := internal.GenerateResponse(`generate for me a "` + message + `?"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	fmt.Println("The result has been copied to the clipboard.")
}
func init() {
	reformulateCmd := &cobra.Command{
		Use:     "r",
		Version: "0.1.0",
		Short:   "Reformulate a given text",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				println("The message to be reformulated cannot be empty.")
			} else {
				reformulate(internal.BuildMessage(args))
			}
		},
	}
	correctCmd := &cobra.Command{
		Use:     "c",
		Version: "0.1.0",
		Short:   "Correct a given text",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				println("The message to be corrected cannot be empty.")
			} else {
				correct(internal.BuildMessage(args))
			}
		},
	}
	answerCmd := &cobra.Command{
		Use:     "a",
		Version: "0.1.0",
		Short:   "Answer a given question",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				println("The question to be answered cannot be empty.")
			} else {
				ask(internal.BuildMessage(args))
			}
		},
	}
	generateCmd := &cobra.Command{
		Use:     "g",
		Version: "0.1.0",
		Short:   "Generate a given request",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				println("The question to be answered cannot be empty.")
			} else {
				generate(internal.BuildMessage(args))
			}
		},
	}
	Cmd.AddCommand(reformulateCmd, correctCmd, answerCmd, generateCmd)
}
