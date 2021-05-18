package main

import "github.com/haobird/logger"

func main() {
	logger.Debug("xxx")
	logger.InitWithConfig(logger.DefaultDefinition())
	logger.Debug("uuu")
}
