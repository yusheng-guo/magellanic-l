package config

type Configuration struct {
	Database Database `mapstructure:"database" yaml:"database" json:"database"`
	Log      Log      `mapstructure:"log" yaml:"log" json:"log"`
	Jwt      Jwt      `mapstructure:"jwt" yaml:"jwt" json:"jwt"`
}
