package config

import "github.com/amnestia/xyz-multifinance/internal/database"

// Config config
type Config struct {
	App    string `yaml:"name"`
	Server Server `yaml:"server"`

	Environment string          `json:"environment"`
	Database    database.Config `json:"database"`
	Crypt       Crypt           `json:"crypt"`
	Auth        Auth            `json:"auth"`
}

// Server server config
type Server struct {
	Port string `yaml:"port"`
	Logs struct {
		Info  string `yaml:"info"`
		Error string `yaml:"error"`
	} `yaml:"logs"`
	Timeout int64 `yaml:"timeout"`
	Limiter struct {
		Rate       int64 `yaml:"rate"`
		Burst      int   `yaml:"burst"`
		Expiration int64 `yaml:"expiration"`
	}
	Origin []string `yaml:"origin"`
}

// Auth auth related config
type Auth struct {
	PubKey   string `json:"pub_key"`
	PrivKey  string `json:"priv_key"`
	LocalKey string `json:"local_key"`
	Pepper   string `json:"pepper"`
}

// Crypt encryption related config
type Crypt struct {
	AESKey string `json:"aes_key"`
	HMAC   HMAC   `json:"hmac"`
}

// HMAC HMAC related config
type HMAC struct {
	Key    string `json:"key"`
	Pepper string `json:"pepper"`
}
