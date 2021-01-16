package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/ali-furkqn/stona/auth"
	"github.com/ali-furkqn/stona/tools"
	"github.com/ali-furkqn/stona/tools/logger"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	FirebaseConfig   *firebaseConfig
	Token            string
	RootPath         string
	StoragePath      string
	ImgMaxResolution int
}

var Config = new(ConfigStruct)

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
		Config.ImgMaxResolution = maxRes
	}

	Config.RootPath = GetEnv("ROOT_PATH", "/")
	Config.StoragePath = GetEnv("STORAGE_PATH", "/")
	token := os.Getenv("TOKEN")

	if token == "" {
		key := GetEnv("KEY", tools.RandomKey(64))
		token = auth.Service().GenerateToken(key)
	}

	Config.Token = token
	auth.Service().SetToken(token)

	switch os.Getenv("TYPE") {
	case "firebase":
		{
			logger.Debug("Config", "Storage is set to firebase")
			Config.firebaseConfigLoad()
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
