package config

type DatabaseCfg struct {
	DriverName             string `yaml:"driver_name"`
	DataSource             string `yaml:"data_source"`
	MaxOpenConns           int    `yaml:"max_open_conns"`
	MaxIdleConns           int    `yaml:"max_idle_conns"`
	ConnMaxLifeTimeSeconds int64  `yaml:"conn_max_life_time_seconds"`
	MigrationConnURL       string `yaml:"migration_conn_url"`
	IsDevMode              bool   `yaml:"is_dev_mode"`
}

type RedisCfg struct {
	CacheTimeSeconds    int    `yaml:"cache_time_seconds"`
	ConnectionURL       string `yaml:"connection_url"`
	PoolSize            int    `yaml:"pool_size"`
	DialTimeoutSeconds  int    `yaml:"dial_timeout_seconds"`
	ReadTimeoutSeconds  int    `yaml:"read_timeout_seconds"`
	WriteTimeoutSeconds int    `yaml:"write_timeout_seconds"`
	IdleTimeoutSeconds  int    `yaml:"idle_timeout_seconds"`
}

type CorsCfg struct {
	AllowOrigins []string `yaml:"allow_origins"`
	AllowHeaders []string `yaml:"allow_headers"`
}
