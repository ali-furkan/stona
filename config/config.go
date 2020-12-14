package config

import (
	"os"
	"stona/auth"
	"stona/tools"
	"stona/tools/logger"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type firebaseConfig struct {
	Type                string `json:"type"`
	StorageBucket       string `json:"storage_bucket"`
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
	FirebaseConfig   *firebaseConfig
	Token            string
	RootPath         string
	StoragePath      string
	ImgMaxResolution int
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
		if d == "" {
			logger.Error("Config", f+" field not found")
			break
		}
		opt[f] = d
	}

	c.FirebaseConfig = &firebaseConfig{
		Type:                opt["FB_TYPE"],
		StorageBucket:       opt["FB_STORAGE_BUCKET"],
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

func GetEnv(key string, def string) string {
	d := os.Getenv(key)
	if d == "" {
		d = def
	}
	return d
}

func init() {
	logger.Debug("Config", "Parsing your configuration")

	if err := godotenv.Load(".env"); err != nil {
		logger.Error("Config", err.Error())
	}

	maxResEnv := GetEnv("IMG_MAX_RESOLUTION", "2048")
	if maxRes, err := strconv.Atoi(maxResEnv); err != nil {
		logger.Error("Config Validation", "'IMG_MAX_RESOLUTION' is not valid. Use ascii integer")
	} else {
		config.ImgMaxResolution = maxRes
	}

	config.RootPath = GetEnv("ROOT_PATH", "/")
	config.StoragePath = GetEnv("STORAGE_PATH", "/")
	token := os.Getenv("TOKEN")

	if token == "" {
		key := GetEnv("KEY", tools.RandomKey(64))
		token = auth.Service().GenerateToken(key)
	}

	config.Token = token
	auth.Service().SetToken(token)

	switch os.Getenv("TYPE") {
	case "firebase":
		{
			logger.Debug("Config", "Storage is set to firebase")
			config.firebaseConfigLoad()
			break
		}
	default:
		{
			logger.Error("Config", "Storage type not found")
		}
	}

	logger.Debug("Config", "Your Configuration Loaded")

	logger.Log("Authentication", "Access Token: "+token)
}
