package clients

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/juanfer2/shopy_dc_golang/src/shared/utilities"
	"gopkg.in/yaml.v3"
)

// postgres://postgres:1234@localhost:5432/golang-clean?sslmode=disable

type B2chatConfig struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func CreateConfigClient() *B2chatConfig {
	configClient := B2chatConfig{}
	root := utilities.GetRootFolder()
	enviroment := strings.ToLower(utilities.GetEnv("MODE"))

	yamlFile, err := ioutil.ReadFile(root +
		"/src/shared/infrastructure/clients/b2chat/config/b2chat_" + enviroment + ".yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &configClient)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &configClient
}

func NewB2chatConfigClient(
	url string,
	username string,
	password string,
) *B2chatConfig {
	return &B2chatConfig{
		Url:      url,
		Username: username,
		Password: password,
	}
}
