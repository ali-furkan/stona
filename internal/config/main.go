package config

// ConfigStruct is struct used to configure when creating the Stona server
// Also provides to create an interface to configuration file (it's yaml now)
type ConfigStruct struct {
	UI bool `yaml:"ui"`
	// Prefix for every storage in Stona
	RootPath string `yaml:"root_path"`

	// Configuration of connection
	Connection struct {
		// Listen Address
		Address string `yaml:"address"`
		// Listen Port
		Port int `yaml:"port"`
		// Connection timeout before connection closed (optional)
		Timeout int `yaml:"timeout"`
	}

	// Provides verification of the connection when connecting to the Stona server (optional)
	Authentication struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	// For TLS Connection (optional)
	secure bool `yaml:"secure"`
	TLS    struct {
		Cert string `yaml:"cert"`
		Key  string `yaml:"key"`
	}
	Buckets []*BaseBucketConfig
}

var config *ConfigStruct

func Load(path string) *ConfigStruct {
	config = configLoad(path)

	return config
}

func Config() *ConfigStruct {
	if config == nil {
		Load("")
	}
	return config
}
