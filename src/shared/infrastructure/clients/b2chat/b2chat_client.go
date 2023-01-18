package clients

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}

type B2ChatClient struct {
	Username string
	Password string
	BaseUrl  string
}

func NewB2ChatClient(baseUrl string, username string, password string) *B2ChatClient {
	return &B2ChatClient{
		Username: username,
		Password: password,
		BaseUrl:  baseUrl,
	}
}

func (b2chat *B2ChatClient) GetToken() string {
	basicString := b64.StdEncoding.EncodeToString([]byte(b2chat.Username + ":" + b2chat.Password))
	client := http.Client{}
	//url := fmt.Sprintf("%s/oauth/token", b2chat.BaseUrl)
	request, err := http.NewRequest(
		"POST",
		"https://api.b2chat.io/oauth/token?grant_type=client_credentials",
		nil,
	)

	if err != nil {
		fmt.Print(err.Error())
		//os.Exit(1)
	}

	request.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Basic " + basicString},
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		//os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	var responseObject TokenResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject.AccessToken
}

// func (b2chat B2ChatClient) SendBroadcast() string {
// 	return ""
// }

func (b2chat *B2ChatClient) Fetch(method string, url string) []byte {
	client := http.Client{}
	request, err := http.NewRequest(method, b2chat.BaseUrl+"/"+url, nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	request.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + b2chat.GetToken()},
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
