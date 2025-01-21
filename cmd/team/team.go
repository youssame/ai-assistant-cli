package team

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/youssame/assistant-cli/internal"
	"log"
	"net/http"
	"os"
)

var Cmd = &cobra.Command{
	Use:     "team",
	Version: "0.1.0",
	Short:   "Manage my reports",
}

func check() {
	resp, err := http.Get(os.Getenv("ASSISTANT_DB_HOST"))
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error while checking the elasticsearch server.")
	}
	if resp.StatusCode != 200 {
		log.Fatal("Error while trying to connect with the elasticsearch server. Code=" + string(rune(resp.StatusCode)))
	}
	internal.SuccessAlert()
}

func init() {
	health := &cobra.Command{
		Use:     "c",
		Version: "0.1.0",
		Short:   "Check the elasticsearch health",
		Run: func(cmd *cobra.Command, args []string) {
			check()
		},
	}
	Cmd.AddCommand(health)
}
