package internal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func GenerateResponse(prompt string) (string, error) {
	url := os.Getenv("LLM_HOST")
	prompt = strings.ReplaceAll(prompt, "\"", "\\\"")
	payload := `{
		"model": "` + os.Getenv("LLM_MODEL") + `",
		"prompt": "Using simple, easy, clear, and direct, and short words, and do not include phrases like 'Here's the answer' or any similar introductions, and with no markdown format, ` + prompt + `"
	}`
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Error creating request")
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Error making request")
	}
	defer resp.Body.Close()

	// Read response body line by line
	scanner := bufio.NewScanner(resp.Body)
	var txt string
	for scanner.Scan() {
		line := scanner.Text()
		var jsonLine map[string]interface{}
		if err := json.Unmarshal([]byte(line), &jsonLine); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		if response, ok := jsonLine["response"]; ok {
			txt += fmt.Sprintf(response.(string))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return "", errors.New("Error reading response body")
	}
	return txt, nil
}
