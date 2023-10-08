package data

import "fmt"

//TODO add function call
//TODO add functions argument
type ChatCompletionBody struct {
	Messages         []ChatCompletionMessage `json:"messages"`
	Model            string                  `json:"model"`
	FrequencyPenalty string                  `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int32        `json:"logit_bias,omitempty"`
	MaxTokens        int64                   `json:"max_tokens,omitempty"`
	N                int64                   `json:"n,omitempty"`
	PresencePenalty  float32                 `json:"presence_penalty,omitempty"`
	Stop             []string                `json:"stop,omitempty"`
	Stream           *bool                   `json:"stream,omitempty"`
	Temperature      float32                 `json:"temperature,omitempty"`
	TopP             float32                 `json:"top_p,omitempty"`
	User             string                  `json:"user,omitempty"`
}

// message type for messages argument in chat completions api
type ChatCompletionMessage struct {
	Content       string          `json:"content"`
	FunctionCalle *OpenAIFuncCall `json:"function_call,omitempty"`
	Name          *string         `json:"name,omitempty"`
	Role          string          `json:"role"`
}

// function call type
type OpenAIFuncCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// choice message type
type ChoiceMessage struct {
	Role         string         `json:"role"`
	Content      string         `json:"content"`
	FunctionCall OpenAIFuncCall `json:"function_call"`
}

// response choice type
type ChatResponseChoices struct {
	Index        int64         `json:"index"`
	Message      ChoiceMessage `json:"message"`
	FinishReason string        `json:"finish_reason"`
}

// response usage type
type ChatResponseUsage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

// chat completion response type
type ChatCompletionResponse struct {
	ID      string                `json:"id"`
	Object  string                `json:"object"`
	Created int64                 `json:"created"`
	Model   string                `json:"model"`
	Choices []ChatResponseChoices `json:"choices"`
	Usage   ChatResponseUsage     `json:"usage"`
	Error   *OpenAIError          `json:"error"`
}

// open ai error type
type OpenAIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Error   string `json:"error"`
	Code    string `json:"code"`
}

// returns formatted string of openai error
func (err *OpenAIError) String() string {
	return fmt.Sprintf("message: %s\ntype: %s\nerror: %s\ncode: %s\n", err.Message, err.Type, err.Error, err.Code)
}
