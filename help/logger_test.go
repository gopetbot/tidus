package help

import (
	"github.com/sirupsen/logrus"
	"testing"
)

const (
	help = "help"
)

func LoggerInit(*testing.T) {

	logger := NewLog()
	logger.WithFields(logrus.Fields{
		"Fields": help,
	}).Info("First log...")

}
