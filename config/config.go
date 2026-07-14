package config

import "github.com/moeghifar/libgo/pkg/envy"

type Config struct {
	// AppVersion is injected from the linker at runtime. It has no env tag,
	// so envy skips it.
	AppVersion string

	App struct {
		Name string `env:"APP_NAME" default:"myapp"`
		Env  string `env:"APP_ENV" default:"local"`
		Port int    `env:"APP_PORT" default:"8080"`
	}

	Database struct {
		DSN         string `env:"DB_DSN" required:"true"`
		MaxPoolSize int    `env:"DB_MAX_POOL" default:"10"`
	}

	Log struct {
		Level  string `env:"LOG_LEVEL" default:"info"`
		Format string `env:"LOG_FORMAT" default:"json"`
	}

	AllowedHosts []string `env:"ALLOWED_HOSTS" default:"localhost,127.0.0.1"`
}

func Load() Config {
	var c Config
	if err := envy.Load(&c); err != nil {
		panic("config: " + err.Error())
	}
	return c
}
