package myLogger

import (
	"encoding/json"
	"io/ioutil"

	"go.uber.org/zap"
)

func InitialConfig() {
	configJSON, err := ioutil.ReadFile("pkg/logger/zapConfigs.json")
	if err != nil {
		panic(err)
	}
	var logConfig zap.Config
	if err := json.Unmarshal(configJSON, &logConfig); err != nil {
		panic(err)
	}

	logger, _ := logConfig.Build()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}
