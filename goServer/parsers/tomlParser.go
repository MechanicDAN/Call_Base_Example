package parsers

import (
	"github.com/spf13/viper"
	"log"
)

func GetConfig(configPath string) {

	if configPath != "" {
		viper.SetConfigFile(configPath)
		err := viper.ReadInConfig()
		if err != nil {
			log.Printf("parsers.tomlParser.GetConfig: Unable to read config file: %s", err)
		}
	} else {
		log.Printf("parsers.tomlParser.GetConfig: Config file is not specified.")
	}
}

func GetPort() string {
	return viper.GetString("port")
}

func GetHost() string {
	return viper.GetString("host")
}

func GetDbAddr() string {
	return viper.GetString("DbAddr")
}

func GetDbUser() string {
	return viper.GetString("DbUser")
}

func GetDbPassword() string {
	return viper.GetString("DbPassword")
}

func GetDbDatabase() string {
	return viper.GetString("DbDatabase")
}


