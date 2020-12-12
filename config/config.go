package config

import (
	"fmt"
	"log"
	"os"
	"stona/auth"
	"stona/tools"
	"strings"

	"github.com/joho/godotenv"
)

type firebaseConfig struct {
	Type                string `json:"type"`
	ProjectID           string `json:"project_id"`
	PrivateKeyID        string `json:"private_key_id"`
	PrivateKey          string `json:"private_key"`
	ClientEmail         string `json:"client_email"`
	ClientID            string `json:"client_id"`
	AuthURI             string `json:"auth_uri"`
	TokenURI            string `json:"token_uri"`
	AuthProviderCertURL string `json:"auth_provider_x509_cert_url"`
	ClientCertURL       string `json:"client_x509_cert_url"`
}

type ConfigStruct struct {
	FirebaseConfig *firebaseConfig
	Token          string
	RootPath       string
	StoragePath    string
}

var config = new(ConfigStruct)

func Config() *ConfigStruct {
	return config
}

func (c *ConfigStruct) firebaseConfigLoad() {
	fields := []string{"FB_TYPE", "FB_PROJECT_ID", "FB_PRIVATE_KEY_ID", "FB_PRIVATE_KEY",
		"FB_CLIENT_EMAIL", "FB_CLIENT_ID", "FB_AUTH_URI", "FB_TOKEN_URI",
		"FB_AUTH_PROVIDER_CERT_URL", "FB_CLIENT_CERT_URL", "FB_STORAGE_BUCKET"}

	opt := make(map[string]string)

	for _, f := range fields {
		d := os.Getenv(f)
		fmt.Println(f, d)
		if d == "" {
			log.Fatalf("Config Error: %s field not found", f)
			break
		}
		opt[f] = d
	}

	c.FirebaseConfig = &firebaseConfig{
		Type:                opt["FB_TYPE"],
		ProjectID:           opt["FB_PROJECT_ID"],
		PrivateKeyID:        opt["FB_PRIVATE_KEY_ID"],
		PrivateKey:          opt["FB_PRIVATE_KEY"],
		ClientEmail:         opt["FB_CLIENT_EMAIL"],
		ClientID:            opt["FB_CLIENT_ID"],
		AuthURI:             opt["FB_AUTH_URI"],
		TokenURI:            opt["FB_TOKEN_URI"],
		AuthProviderCertURL: opt["FB_AUTH_PROVIDER_CERT_URL"],
		ClientCertURL:       opt["FB_CLIENT_CERT_URL"],
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	rootPath := os.Getenv("ROOT_PATH")
	storagePath := os.Getenv("STORAGE_PATH")
	token := os.Getenv("TOKEN")

	if rootPath == "" {
		rootPath = "/"
	}
	if storagePath == "" {
		storagePath = "/"
	}
	if token == "" {
		key := os.Getenv("KEY")
		if key == "" {
			key = tools.RandomKey(64)
		}
		token = auth.Service().GenerateToken(key)
	}

	switch os.Getenv("TYPE") {
	case "firebase":
		{
			config.firebaseConfigLoad()
		}
	default:
		{
			log.Fatalln("Config Error: Storage type not found")
		}
	}

}
