package internal

import (
	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

func Copy(text string) {
	clipboard.Write(clipboard.FmtText, []byte(text))
}

func BuildMessage(words []string) string {
	txt := ""
	for i := 0; i < len(words); i++ {
		txt += words[i] + " "
	}
	return txt
}
