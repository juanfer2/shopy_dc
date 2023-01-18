package postgres

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/juanfer2/shopy_dc_golang/src/shared/utilities"
	"gopkg.in/yaml.v3"
)

// postgres://postgres:1234@localhost:5432/golang-clean?sslmode=disable

type ConfigClient struct {
	Drive           string `yaml:"drive"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	Database        string `yaml:"database"`
	MigrationFolder string
}

func CreateConfigClient() *ConfigClient {
	configClient := ConfigClient{}
	root := utilities.GetRootFolder()
	enviroment := strings.ToLower(utilities.GetEnv("MODE"))

	yamlFile, err := ioutil.ReadFile(root +
		"/src/shared/infrastructure/persistence/postgres/config/database_" + enviroment + ".yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &configClient)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	setMigrationFolder(&configClient)

	return &configClient
}

func NewConfigClient(
	drive string,
	username string,
	password string,
	host string,
	port string,
	database string,
) *ConfigClient {
	return &ConfigClient{
		Drive:           drive,
		Username:        username,
		Password:        password,
		Host:            host,
		Port:            port,
		Database:        database,
		MigrationFolder: migrationFolder(),
	}
}

func setMigrationFolder(configClient *ConfigClient) {
	configClient.MigrationFolder = migrationFolder()
}

func (createConfigClient ConfigClient) GetUrl() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		createConfigClient.Drive, createConfigClient.Username, createConfigClient.Password,
		createConfigClient.Host, createConfigClient.Port, createConfigClient.Database,
	)
}

func migrationFolder() string {
	root := utilities.GetRootFolder()

	return "file://" + root + "/src/shared/infrastructure/persistence/postgres/migrations"
}
