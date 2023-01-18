package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Gpt3Client struct {
	BaseUrl string
	token   string
}

type Gpt3FetchParams struct {
	Method string
	Url    string
	Body   any
}

func NewGpt3Client(baseUrl string, token string) Gpt3Client {
	return Gpt3Client{
		BaseUrl: baseUrl,
		token:   token,
	}
}

func (gpt3Client *Gpt3Client) GetModels() []Gpt3Model {
	response, err := gpt3Client.fetch(Gpt3FetchParams{Method: GET, Url: MODELS})
	if err != nil {
		fmt.Print(err.Error())
		//os.Exit(1)
	}

	return Gpt3ModelsFromResponse(response)
}

func (gpt3Client *Gpt3Client) Completion(gpt3CompletionRequest Gpt3CompletionRequest) Gpt3Complention {
	response, err := gpt3Client.fetch(Gpt3FetchParams{Method: http.MethodPost, Url: COMPLETIONS,
		Body: gpt3CompletionRequest})
	if err != nil {
		fmt.Print(err.Error())
		//os.Exit(1)
	}

	return NewGpt3ComplentionFromResponse(response)
}

func (gpt3Client *Gpt3Client) fetch(gpt3FetchParams Gpt3FetchParams) ([]byte, error) {
	var bodyRequest []byte = nil
	if gpt3FetchParams.Body != nil {
		jsonReq, err := json.Marshal(gpt3FetchParams.Body)
		if err != nil {
			log.Fatalln(err)
		}

		bodyRequest = jsonReq
	}

	client := http.Client{}
	request, err := http.NewRequest(
		gpt3FetchParams.Method,                        //"POST",
		gpt3Client.BaseUrl+"/v1/"+gpt3FetchParams.Url, //"https://api.b2chat.io/oauth/token?grant_type=client_credentials",
		bytes.NewBuffer(bodyRequest),
	)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
		//os.Exit(1)
	}

	request.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + gpt3Client.token},
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
		//os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var anyResponse any
	json.Unmarshal(responseData, &anyResponse)

	fmt.Println(anyResponse)

	return responseData, err
}
