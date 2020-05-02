package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

var Logz *zap.Logger

func init() {
	rawJSON := []byte(`{
      "level": "info",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stderr"]
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	var err error
	Logz, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	defer func() {

	}()

	Logz.Info("logger construction succeeded")
}
