package configurations

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	ConsumerKey    string `json:"consumerkey"`
	ConsumerSecret string `json:"consumer_secret"`
	CallbackURL    string `json:"calbackUrl"`
	ServeAddr      string `json:"serveAddr"`
}

func GetConfiguration() *Configuration {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	var config Configuration

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(byteValue, &config)

	return &config
}
