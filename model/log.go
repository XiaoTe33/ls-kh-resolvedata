package model

import (
	"github.com/sirupsen/logrus"
	"ls-kh-resolvedata/log"
)

var myLog = log.Log

func init() {
	myLog.SetLevel(logrus.TraceLevel)
}
