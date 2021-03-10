package config

type GCSConfig struct {
	Type                string `yaml:"type"`
	StorageBucket       string `yaml:"storage_bucket"`
	ProjectID           string `yaml:"project_id"`
	PrivateKeyID        string `yaml:"private_key_id"`
	PrivateKey          string `yaml:"private_key"`
	ClientEmail         string `yaml:"client_email"`
	ClientID            string `yaml:"client_id"`
	AuthURI             string `yaml:"auth_uri"`
	TokenURI            string `yaml:"token_uri"`
	AuthProviderCertURL string `yaml:"auth_provider_x509_cert_url"`
	ClientCertURL       string `yaml:"client_x509_cert_url"`
}
