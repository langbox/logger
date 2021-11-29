package main

import "github.com/langbox/logger"

func main() {
	logger.Debug("xxx")
	logger.InitWithConfig(logger.DefaultDefinition())
	logger.Debug("uuu")
}
