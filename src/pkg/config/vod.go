package config

type VodConfig struct {
	// 是否开启Vod，默认关闭
	Enabled bool `json:"enable" mapstructure:"enable"`

	OpenAPIConfig
}
