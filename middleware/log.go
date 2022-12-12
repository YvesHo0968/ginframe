package middleware

import "github.com/sirupsen/logrus"

// Logger 日志到文件
func Logger() {
	var logger *logrus.Logger
	logger = logrus.New()

	logger.SetLevel(logrus.DebugLevel)
}
