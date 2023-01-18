package clients

import "encoding/json"

type Gpt3CompletionRequest struct {
	Model       string   `json:"model"`
	Prompt      []string `json:"prompt"`
	Suffix      string   `json:"suffix"`
	MaxTokens   string   `json:"max_tokens"`
	Temperature string   `json:"temperature"`
	N           uint     `json:"n"`
	Echo        bool     `json:"echo"`
}

type Gpt3Usage struct {
	PromptTokens     uint `json:"prompt_tokens"`
	CompletionTokens uint `json:"completion_tokens"`
	TotalTokens      uint `json:"total_tokens"`
}

type Gpt3ComplentionChoices struct {
	Text         string `json:"text"`
	Index        uint   `json:"index"`
	Logprobs     string `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type Gpt3Complention struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created string                 `json:"created"`
	Choices Gpt3ComplentionChoices `json:"choices"`
	Usage   Gpt3Usage              `json:"usage"`
}

func NewGpt3ComplentionFromResponse(response []byte) (gpt3Complention Gpt3Complention) {
	json.Unmarshal(response, &gpt3Complention)

	return gpt3Complention
}
