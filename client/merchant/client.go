package merchant

import (
	"github.com/Dexus/go-paypal-classic/client"
	"github.com/Dexus/go-paypal-classic/config"
)

// New creates client.Client with given config
func New(conf *config.Config) client.Client {
	return client.Client{
		Config: conf,
	}
}
