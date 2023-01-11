package startup

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString string
}

var AppConfig Config

func loadAppConfig() {
	log.Println("Loading server configuration")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	AppConfig.ConnectionString = viper.GetString("MONGODB_CONNECTION_STRING")
}

func init() {
	loadAppConfig()
}
