package cmd

import (
	"cat-project/internal/logging"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func Start() {
	fileName, _:= filepath.Abs("./config/config.yaml")
	yamlBytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	m := map[string] interface{}{}
	yaml.Unmarshal(yamlBytes, m)
	logMap := m["logging"].(map[string] interface{})
	logLevel := logMap["level"].(int)
	logger := logging.New(logging.Config{LogLevel: logLevel})
	logger.Info("Message Hello World")
}