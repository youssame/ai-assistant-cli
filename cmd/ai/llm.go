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
	res, err := internal.GenerateResponse(`reformulate this message in a simple easy words, "` + message + `"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	internal.ClipboardSuccess()
}

// correct a given text
func correct(message string) {
	res, err := internal.GenerateResponse(`correct this "` + message + `"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	internal.ClipboardSuccess()
}

// ask a given question
func ask(message string) {
	res, err := internal.GenerateResponse(`answer this question "` + message + `?"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	internal.ClipboardSuccess()
}

// genEmail an email based on a given topic
func genEmail(message string) {
	res, err := internal.GenerateResponse(`generate a email about : "` + message + `"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	internal.ClipboardSuccess()
}

// genEmail an email based on a given topic
func genMessage(message string) {
	res, err := internal.GenerateResponse(` generate a slack message about "` + message + `"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	internal.ClipboardSuccess()
}

// genNotes generates a professional and organized meeting notes from a draft
func genNotes(message string) {
	res, err := internal.GenerateResponse(` generate meeting notes. Start the notes with 'Date: [today's date]'. After that, use bullet points to summarize the key points discussed in the meeting (the notes are seperated by "/") : "` + message + `"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	internal.ClipboardSuccess()
}

// genAnnouncement generates a announcement
func genAnnouncement(message string) {
	res, err := internal.GenerateResponse(` Generate an announcement with a title starting with an emoji and a description (please start the description with '@channel') about : "` + message + `"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	internal.ClipboardSuccess()
}

// generate a given request (announcement, email, message...)
func generate(message string) {
	res, err := internal.GenerateResponse(`generate for me a "` + message + `?"`)
	if err != nil {
		log.Fatal(err)
	}
	internal.Copy(res)
	internal.ClipboardSuccess()
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
	emailCmd := &cobra.Command{
		Use:     "e",
		Version: "0.1.0",
		Short:   "Generate an email based on a given topic",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				println("The message to be corrected cannot be empty.")
			} else {
				genEmail(internal.BuildMessage(args))
			}
		},
	}
	msgCmd := &cobra.Command{
		Use:     "m",
		Version: "0.1.0",
		Short:   "Generate a slack message based on a given topic",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				println("The message to be corrected cannot be empty.")
			} else {
				genMessage(internal.BuildMessage(args))
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
	genNotesCmd := &cobra.Command{
		Use:     "n",
		Version: "0.1.0",
		Short:   "Generate Meeting Notes from a Draft",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("The question to be answered cannot be empty.")
			} else {
				genNotes(internal.BuildMessage(args))
			}
		},
	}
	genAnnouncementCmd := &cobra.Command{
		Use:     "an",
		Version: "0.1.0",
		Short:   "Generate an announcement from a Draft",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("The question to be answered cannot be empty.")
			} else {
				genAnnouncement(internal.BuildMessage(args))
			}
		},
	}
	Cmd.AddCommand(reformulateCmd, correctCmd, answerCmd, generateCmd, emailCmd, msgCmd, genNotesCmd, genAnnouncementCmd)
}
