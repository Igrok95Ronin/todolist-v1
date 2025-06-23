package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

// Структура конфигурации
type Config struct {
	Port string `yaml:"port"`
}

// Глобальная переменная для хранения конфигурации
var (
	configInstance *Config   // Синглтон-объект (единственный экземпляр)
	once           sync.Once // Гарантирует однократное выполнение
)

// Функция получения конфигурации
func GetConfig() *Config {
	// Загружаем .env
	//if err := godotenv.Load(); err != nil {
	//	fmt.Println("Не удалось загрузить .env файл, переменные окружения могут отсутствовать")
	//}

	once.Do(func() {
		configInstance = &Config{}

		if err := cleanenv.ReadConfig("./config.yml", configInstance); err != nil {
			help, _ := cleanenv.GetDescription(configInstance, nil)
			fmt.Println(help)
		}

		// Читаем переменные окружения и заменяем значения в конфиге
		//overrideWithEnv(configInstance)
	})

	return configInstance
}

// overrideWithEnv перезаписывает конфигурацию значениями из переменных окружения (если они заданы)
//func overrideWithEnv(cfg *Config)  {
//	if user := os.Getenv("POSTGRES_USER"); user != "" {
//		cfg.DB.User = user
//	}
//	if password := os.Getenv("POSTGRES_PASSWORD"); password != "" {
//		cfg.DB.Password = password
//	}
//	if dbName := os.Getenv("POSTGRES_DB"); dbName != "" {
//		cfg.DB.DBName = dbName
//	}
//	if accessToken := os.Getenv("ACCESS_TOKEN"); accessToken != "" {
//		cfg.Token.Access = accessToken
//	}
//	if refreshToken := os.Getenv("REFRESH_TOKEN"); refreshToken != "" {
//		cfg.Token.Refresh = refreshToken
//	}
//}
