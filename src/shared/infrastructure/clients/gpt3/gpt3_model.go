package clients

import "encoding/json"

type Gpt3Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}

type Gpt3ModelResponse struct {
	Data []Gpt3Model `json:"data"`
}

func Gpt3ModelsFromResponse(response []byte) []Gpt3Model {
	var gpt3ModelResponse Gpt3ModelResponse
	json.Unmarshal(response, &gpt3ModelResponse)

	return gpt3ModelResponse.Data
}
