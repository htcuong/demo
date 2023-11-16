package config

type Configurations struct {
	LocalDevelopment bool      `mapstructure:"localDevelopment"`
	Database         *Database `mapstructure:"database"`
	Service          Service   `mapstructure:"service"`
}

type Service struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Hostname string `mapstructure:"hostname"`
	DBname   string `mapstructure:"dbname"`
}
