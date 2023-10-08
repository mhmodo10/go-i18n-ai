package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mhmodo10/go-i18n-ai/internal/data"
)

// Open ai client type
type OpenAIClient struct {
	APIKey             string
	ChatCompletionsURL string
}

// Creates new client and returns pointer to it
func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		APIKey:             apiKey,
		ChatCompletionsURL: "https://api.openai.com/v1/chat/completions",
	}
}

// Calls chatcompletion endpoint
func (oAI *OpenAIClient) ChatCompletion(chat data.ChatCompletionBody) (*data.ChatCompletionResponse, error) {
	responseBody := data.ChatCompletionResponse{}
	body, err := json.Marshal(chat)
	if err != nil {
		return &responseBody, err
	}
	req, err := http.NewRequest(http.MethodPost, oAI.ChatCompletionsURL, bytes.NewBuffer(body))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oAI.APIKey))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return &responseBody, err
	}
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return &responseBody, err
	}
	err = json.NewDecoder(res.Body).Decode(&responseBody)
	if err != nil {
		return &responseBody, err
	}
	return &responseBody, nil
}
