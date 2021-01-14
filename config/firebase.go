package config

import (
	"os"

	"github.com/ali-furkqn/stona/tools/logger"
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
