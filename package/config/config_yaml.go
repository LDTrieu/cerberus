package config

type DatabaseCfg struct {
	IsDevMode              bool   `yaml:"is_dev_mode"`
}

type RedisCfg struct {
	IdleTimeoutSeconds  int    `yaml:"idle_timeout_seconds"`
}

type CorsCfg struct {
	AllowHeaders []string `yaml:"allow_headers"`
}
