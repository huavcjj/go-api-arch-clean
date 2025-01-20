package database

import "github.com/caarlos0/env/v11"

type ConfigMySQL struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Database string `env:"DB_NAME" envDefault:"api_database"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	Driver   string `env:"DB_DRIVER" envDefault:"mysql"`
	User     string `env:"DB_USER" envDefault:"app"`
	Password string `env:"DB_PASSWORD" envDefault:"password"`
}

type ConfigSQLite struct {
	Database string `env:"DB_NAME" envDefault:"api_database.sqlite"`
}

func NewConfigMySQL() *ConfigMySQL {
	var config ConfigMySQL
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
	return &config
}

func NewConfigSQLite() *ConfigSQLite {
	var config ConfigSQLite
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
	return &config
}

// func NewConfigMySQL() *Config {
// 	return &Config{
// 		Host:     pkg.GetEnvDefault("DB_HOST", "localhost"),
// 		Database: pkg.GetEnvDefault("DB_NAME", "api_database"),
// 		Port:     pkg.GetEnvDefault("DB_PORT", "3306"),
// 		Driver:   pkg.GetEnvDefault("DB_DRIVER", "mysql"),
// 		User:     pkg.GetEnvDefault("DB_USER", "app"),
// 		Password: pkg.GetEnvDefault("DB_PASSWORD", "password"),
// 	}
// }

// func NewConfigSQLite() *Config {
// 	return &Config{
// 		Database: pkg.GetEnvDefault("DB_NAME", "api_database.sqlite"),
// 	}
// }
