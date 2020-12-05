package help

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type Logging struct {
	*log.Logger
}

func (l *Logging) Config() {
	l.SetFormatter(&log.JSONFormatter{})
	l.SetOutput(os.Stdout)
}

func NewLog() *Logging {
	var logging = &Logging{
		log.New(),
	}
	logging.Config()

	return logging
}
