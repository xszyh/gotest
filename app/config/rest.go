package config

type Config struct{
	Port string
}

GetConfig() *Config{
	return &Config{"8011"}
}