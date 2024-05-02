package logger

import (
	"log"
	"mooi/config"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once

var Log *logrus.Logger

func init() {
	once.Do(func() {
		Log = newLogger()
	})
}

func newLogger() *logrus.Logger {
	Log := logrus.New()
	log.Println("Setup Logger")
	if config.Get().LogPath != "" {
		err := os.Mkdir(config.Get().LogPath, 0755)
		if err != nil {
			log.Println("[ERROR] failed to create log path. error: ", err)
		}
		log.Println("Success to create log path")
	}

	// write log output to file
	logFilename := "server.log"
	logfile, err := os.OpenFile(logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println("[ERROR] not able to open/create log file. err: ", err)
	}
	Log.SetOutput(logfile)

	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	Log.SetReportCaller(true)

	return Log
}
