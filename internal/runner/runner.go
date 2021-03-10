package runner

import (
	"github.com/ali-furkqn/stona/internal/config"
	"github.com/ali-furkqn/stona/internal/transport"
)

func Start(configPath string) {
	// Load Config
	config.Load(configPath)

	// Start the tcp transport
	transport.Serve(&transport.ServerConfig{
		Address: config.Config().Connection.Address,
		Port:    config.Config().Connection.Port,
		TLS: struct {
			Cert string
			Key  string
		}(config.Config().TLS),
	})
}
