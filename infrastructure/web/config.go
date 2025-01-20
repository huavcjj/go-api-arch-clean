package web

import "github.com/caarlos0/env/v11"

type Config struct {
	Host             string   `env:"WEB_HOST" envDefault:"0.0.0.0"`
	Port             string   `env:"WEB_PORT" envDefault:"8080"`
	CorsAllowOrigins []string `env:"WEB_CORS_ALLOW_ORIGINS" envDefault:"http://0.0.0.0:8001"`
}

func NewConfigWeb() *Config {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
	return &config
}

// func NewConfigWeb() *Config {
// 	return &Config{
// 		Host: pkg.GetEnvDefault("WEB_HOST", "0.0.0.0"),
// 		Port: pkg.GetEnvDefault("WEB_PORT", "8080"),
// 		CorsAllowOrigins: strings.Split(pkg.GetEnvDefault(
// 			"WEB_CORS_ALLOW_ORIGINS",
// 			"http://0.0.0.0:8001"), ","),
// 	}
// }
