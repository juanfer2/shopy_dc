package clients

func CreateB2ChatClient() *B2ChatClient {
	config := CreateConfigClient()
	return NewB2ChatClient(config.Url, config.Username, config.Password)
}
