package filelog

import (
	"log"
	"os"
	"path"
)

// функция создания фала логирования
func fileLog(loggDir, fname string) *log.Logger {

	if _, err := os.Stat(loggDir); os.IsNotExist(err) {
		os.MkdirAll(loggDir, os.ModePerm)
	}

	p := path.Join(loggDir, fname+".log")
	file, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	logger := log.New(file, fname+": ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
