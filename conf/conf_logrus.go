package conf

type Log struct { // Log - структура для хранения конфигурации логгера
	App string `yaml:"app"` // App - название приложения
	Dir string `yaml:"dir"` // Dir - путь к директории для логов
}
