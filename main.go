package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

const (
	apiURL = "https://api.openai.com/v1/chat/completions"
)

type Message struct {
  Content string `json:"content"`
  Role string `json:"role"`
}

type chatRequest struct {
	Model     string `json:"model"`
	MaxTokens int    `json:"max_tokens"`
  Messages  []Message `json:"messages"`
}

type chatResponse struct {
  Choices []struct {
    Message struct {
      Content string `json:"content"`
      Role string `json:"role"`
    } `json:"message"`
  } `json:"choices"`
}

var apiKey string

func init() {
	flag.StringVar(&apiKey, "api_key", "", "OpenAI API key")
	flag.Parse()

	if apiKey == "" {
		fmt.Fprintln(os.Stderr, color.RedString("Error: API key is required"))
		os.Exit(1)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
  messageHistory := []Message{}

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, color.RedString("Failed to read input: %v\n", err))
			os.Exit(1)
		}

		input = strings.TrimSpace(input)
    messageHistory = append(messageHistory, Message{
      Content: input,
      Role: "user",
    })

		if input == "" {
			continue
		}

		if input == "exit" {
			break
		}

		response, err := chat(messageHistory)
		if err != nil {
			fmt.Fprintf(os.Stderr, color.RedString("Failed to chat: %s\n", err.Error()))
			continue
		}

		response = strings.TrimSpace(response)
    messageHistory = append(messageHistory, Message{
      Content: response,
      Role: "assistant",
    })
		if response != "" {
			fmt.Println(color.GreenString(response))
		}
	}
}

func chat(messageHistory []Message) (string, error) {
	client := &http.Client{}

	reqBody := chatRequest{
		Messages:   messageHistory,
		Model:     "gpt-3.5-turbo",
		MaxTokens: 256,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("Failed to encode request body: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(respBody))
	}

	var chatResp chatResponse
	err = json.Unmarshal(respBody, &chatResp)
	if err != nil {
		return "", fmt.Errorf("Failed to parse response body: %v", err)
	}

	if len(chatResp.Choices) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}

	return "", nil
}

