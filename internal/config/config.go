package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

// Структура конфигурации
type Config struct {
	Port string         `yaml:"port"`
	DB   DatabaseConfig `yaml:"db"`
}

// Подконфигурация для базы данных
type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	SslMode  string `yaml:"sslMode"`
	TimeZone string `yaml:"timeZone"`
}

// Глобальная переменная для хранения конфигурации
var instance *Config
var once sync.Once

// Функция получения конфигурации
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			fmt.Println(help)
		}
	})
	return instance
}
