package configs

import (
	"github.com/spf13/viper"
	"time"
)

type conf struct {
	DBHost                 string        `mapstructure:"DB_HOST"`
	DBPort                 string        `mapstructure:"DB_PORT"`
	DBPassword             string        `mapstructure:"DB_PASSWORD"`
	KeyLimiter             bool          `mapstructure:"KEY_LIMITER"`
	KeyLimit               int           `mapstructure:"KEY_LIMIT"`
	IPLimiter              bool          `mapstructure:"IP_LIMITER"`
	IPLimit                int           `mapstructure:"IP_LIMIT"`
	RequestLimiterDuration time.Duration `mapstructure:"REQUEST_LIMITER_DURATION"`
	RequestBlockerDuration time.Duration `mapstructure:"REQUEST_BLOCKER_DURATION"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
