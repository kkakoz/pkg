package logger_test

import (
	"github.com/kkakoz/pkg/logger"
	"github.com/spf13/viper"
	"testing"
)

func TestLogger(t *testing.T) {
	logger.InitLog(viper.GetViper())
	logger.Info("test logger")
}
