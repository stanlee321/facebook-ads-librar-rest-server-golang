package util

import (
	"github.com/spf13/viper"
)

// Config strores all the configuration of the application
// the values are read by viper from a config file or env var
type Config struct {
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBSource       string `mapstructure:"DB_SOURCE"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	GRPCAddress    string `mapstructure:"GRPC_ADDRESS"`
	GRPCETLAddress string `mapstructure:"GRPC_ETL_ADDRESS"`
}

// LoadConfig reads configuration from file or env var
func LoadConfig(path string, mode string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
