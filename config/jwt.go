package config

import "time"

type Jwt struct {
	Key string        `mapstructure:"key" json:"key" yaml:"key"` // 密钥
	TTL time.Duration `mapstructure:"ttl" json:"ttl" yaml:"ttl"` // Time to Live 寿命
}
