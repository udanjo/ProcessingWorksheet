package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ueverson/ProcessingWorksheetGO/middleware"
)

type Configuration struct {
	UrlPlanilha string `json:"UrlPlanilha"`
}

//Faz a leitura da configuração do config
func Config() Configuration {
	file, err := os.Open("configs/config.json")
	middleware.Handler(err)
	defer file.Close()

	configBytes, err := ioutil.ReadAll(file)
	middleware.Handler(err)

	var config Configuration
	err = json.Unmarshal(configBytes, &config)
	middleware.Handler(err)

	return config
}
