package config

import (
	"time"
)

// Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
	Time              time.Duration
	ServiceName       string
	TimeConvert       string
	JwtCertFile       string
}

// Logger config
type Logger struct {
	DisableCaller     bool   `yaml:"disable_caller"`
	DisableStacktrace bool   `yaml:"disable_stacktrace"`
	Encoding          string `yaml:"encoding"`
	Level             string `yaml:"level"`
}

// MySql config
type MysqlConfig struct {
	MySqlHost     string
	MySqlPort     string
	MySqlUser     string
	MySqlPassword string
	MySqlDbname   string
	MySqlSSLMode  bool
	MysqlDriver   string
	MySqlLogLevel string
}

// Redis config
type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}

// Cookie config
type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

// Session config
type Session struct {
	Prefix string
	Name   string
	Expire int
}

// JwtKey config
type JwtKeyConfig struct {
	SecretKey       string
	AuthPublicKey1  string
	AuthPrivateKey1 string
	AuthPublicKey2  string
	AuthPrivateKey2 string
	TTL             int
}

// Cors config
type CorsConfig struct {
	AllowOrigins []string
	AllowHeaders []string
}

// RabbitMQ
type RabbitMQConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Vhost    string
}

type ConsumerMessage []ExchangeDetail
type PublisherMessage []ExchangeDetail

type ExchangeDetail struct {
	Queue struct {
		Name string
		Size int
	}
	RoutingKey     string
	Exchange       string
	ConsumerTag    string
	WorkerPoolSize int
	DeadLetter     DeadLetter
}
type DeadLetter struct {
	Exchange string
	Queue    string
	RetryMax int
	TTL      int
}

// Api External config
type ApiExternal struct {
	Hasaki      string
	HasakiLogin bool
}

type DiscordConfig struct {
	WebhookUrl string
	User       string
	Avatar     string
}

type KafkaConfig struct {
	Addrs           []string
	Topics          []string
	Group           string
	GroupId         string
	MaxMessageBytes int
	Compress        bool
	Newest          bool
	Version         string

	Consumer struct {
		GroupHeartbeatInterval int64
		GroupSessionTimeout    int32
		MaxProcessingTime      int32
		ReturnErrors           *bool
	}
	Acl struct {
		Enable   bool
		User     string
		Password string
	}
}

type ESConfig struct {
	Addrs    []string
	User     string
	Password string
}

// Get config path by app environment
func GetConfigPath(configPath string) string {
	if configPath == "qc" {
		return "./config/config-qc"
	}
	if configPath == "staging" {
		return "./config/config-staging"
	}
	if configPath == "prod" {
		return "./config/config-prod"
	}
	return "./config/config-dev"
}
