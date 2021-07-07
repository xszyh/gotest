package config

type Config struct {
	Port string
}

func GetConfig() *Config {
	return &Config{"8011"}
}
