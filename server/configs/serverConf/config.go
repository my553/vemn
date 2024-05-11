package serverConf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	POSTGRESQL_CONNSTRING string `mapstructure:"POSTGRESQL_CONNSTRING"`
}

var DefaultConfig Config

// LoadConfig функция обработки конфига
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintln("error in config: $v", err))
	}

	err = viper.Unmarshal(&config)
	DefaultConfig = config
	return
}
