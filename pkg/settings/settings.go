package settings

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// сборка информации о логах
type Settings struct {
	ServiceName string `env:"SERVICE_NAME"`
	LoggerType  string `env:"LOG_TYPE"`
	LogLevel    string `env:"LOG_LEVEL"`  //уровень логгирования

	SysLog   SysLog
	FileLog  FileLog
	EmailLog EmailLog
}

type SysLog struct {
	Host string `env:"LOG_SERVER_HOST"` //Хост сервера логов
	Port string `env:"LOG_SERVER_PORT"` //Порт сервера логов

	Facility int `env:"FACILITY"`
}
type FileLog struct {
	NameLogger string `env:"NAME_LOGGER"` //имя папки хранения логгера
	LogDir     string `env:"LOG_DIR"`     //директория хранения
}
type EmailLog struct {
	From     string `env:"FROM"`     //информация об отправителе(почта)
	SmtpHost string `env:"SMTPHOST"` //smtp хост
	SmtpPort string `env:"SMTPPORT"` //smtp порт
	To       string `env:"TO"`       //кому отправляется почта
}

// формирование структуры setting из .env и данных настройки сервера
func GetSettingsEnv() (*Settings, error) {

	var err error
	s := Settings{}
	godotenv.Load()

	err = env.Parse(&s)
	if err != nil {
		log.Fatalf("невозможно проанализировать переменные среды %e", err)
		return nil, err
	}
	return &s, nil
}
