package configurations

import (
	"os"
)

type Configuration struct {
	TwitterToken string
	CallbackURL  string
	HostAddr     string
}

func GetConfiguration() *Configuration {
	return &Configuration{
		TwitterToken: os.Getenv("TWITTER_TOKEN"),
		CallbackURL:  os.Getenv("CALLBACK_URL"),
		HostAddr:     os.Getenv("HOST_ADDRESS"),
	}
}
