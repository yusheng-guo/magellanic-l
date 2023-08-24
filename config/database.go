package config

type Database struct {
	Driver                string `mapstructure:"driver" yaml:"driver" json:"driver"`
	DBName                string `mapstructure:"db_name" yaml:"db_name" json:"db_name"`
	Host                  string `mapstructure:"host" yaml:"host" json:"host"`
	Port                  int    `mapstructure:"port" yaml:"port" json:"port"`
	Username              string `mapstructure:"username" yaml:"username" json:"username"`
	Password              string `mapstructure:"password" yaml:"password" json:"password"`
	Charset               string `mapstructure:"charset" yaml:"charset" json:"charset"`
	Timeout               int    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`
	MaximumConnections    int    `mapstructure:"maximum_connections" yaml:"maximum_connections" json:"maximum_connections"`
	MinimumConnections    int    `mapstructure:"minimum_connections" yaml:"minimum_connections" json:"minimum_connections"`
	IdleConnectionTimeout int    `mapstructure:"idle_connection_timeout" yaml:"idle_connection_timeout" json:"idle_connection_timeout"`
}
