package config

import (
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	LiveStreamConfig LiveStreamConfig `json:"live_stream" mapstructure:"live_stream"`
	LiveMicConfig    LiveMicConfig    `json:"live_mic" mapstructure:"live_mic"`
	LiveIMConfig     LiveIMConfig     `json:"live_im" mapstructure:"live_im"`
	OpenAPIConfig    OpenAPIConfig    `json:"openapi" mapstructure:"openapi"`
	StorageConfig    StorageConfig    `json:"storage" mapstructure:"storage"`

	GatewayConfig GatewayConfig `json:"gateway" mapstructure:"gateway"`
}

type GatewayConfig struct {
	Enable  bool   `json:"enable" mapstructure:"enable"`
	WsAddr  string `json:"ws_addr" mapstructure:"ws_addr"`
	LwpAddr string `json:"lwp_addr" mapstructure:"lwp_addr"`
}

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("./conf") // path to look for the config file in
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	viper.SetEnvPrefix("SPF")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, err
	}
	return &appConfig, nil
}
