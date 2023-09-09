package configurator

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

type RGBA struct {
	R uint8 `json:"red"`
	G uint8 `json:"green"`
	B uint8 `json:"blue"`
	A uint8 `json:"alpha"`
}

type Config struct {
	HistoryLocation string `json:"history_location"`
	LogsLocation    string `json:"logs_location"`
	ButtonsColor    RGBA   `json:"buttons_color"`
	LabelsColor     RGBA   `json:"labels_color"`
	EntriesColor    RGBA   `json:"entries_color"`
}

func New() Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonBody, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err = json.Unmarshal(jsonBody, &config); err != nil {
		logrus.Fatal(err)
	}

	return config
}
