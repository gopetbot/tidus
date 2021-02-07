package help

import (
	"github.com/sirupsen/logrus"
	"testing"
)

const (
	help = "help"
)

func Test_LoggerInit(*testing.T) {

	logger := NewLog()
	logger.WithFields(logrus.Fields{
		"Fields": help,
	}).Info("First log...")

}
