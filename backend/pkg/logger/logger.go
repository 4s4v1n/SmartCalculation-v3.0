package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

type Logger struct {
	file *os.File
}

func New(path string) *Logger {
	if _, err := os.Open(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, 0777); err != nil {
			log.Fatal(err)
		}
	}

	now := time.Now()

	fileName := fmt.Sprintf("logs_%s-%d-%d-%d", now.Format("02-01-2006"),
		now.Hour(), now.Minute(), now.Second())

	file, err := os.Create(path + "/" + fileName + ".log")
	if err != nil {
		log.Fatal(err)
	}

	logrus.SetOutput(file)

	return &Logger{
		file: file,
	}
}

func (l *Logger) Release() {
	if err := l.file.Close(); err != nil {
		log.Fatal()
	}
}
