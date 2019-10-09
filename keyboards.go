// Package go-qmk provides QMK API wrapper utilities
package qmk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const qmkAPI = "https://api.qmk.fm/v1/keyboards"

var httpClient = &http.Client{
	Timeout: time.Second * 2,
}

// A Keyboard contains metadata as to available layouts,
// keymaps, identifier etc...
type Keyboard struct {
	Name          string   `json:"name"`
	URL           string   `json:"url"`
	Keymaps       []string `json:"keymaps"`
	GitHash       string   `json:"git_hash"`
	BootLoader    string   `json:"bootloader"`
	Platform      string   `json:"platform"`
	Processor     string   `json:"processor,omitempty"`
	ProcessorType string   `json:"processor_type"`
	DeviceVersion string   `json:"device_ver"`
	Identifier    string   `json:"identifier"`
	Maintainer    string   `json:"maintainer,omitempty"`
}

// A rawData is the super data structure for Keyboard
type rawData struct {
	LastUpdated string              `json:"last_updated"`
	Keyboards   map[string]Keyboard `json:"keyboards"`
}

// Keyboards queries QMK API and returns an array of keyboard
// names.
func AllKeyboardsList() []string {
	var keyboardsList []string
	keyboardsData := queryQMK("all")

	for k, _ := range keyboardsData {
		keyboardsList = append(keyboardsList, k)
	}

	return keyboardsList
}

// Keyboard queries QMK API and returns a Keyboard containing
// its particulars.
func KeyboardData(keyboard string) Keyboard {
	rawData := queryQMK(keyboard)
	return rawData[keyboard]
}

// Keymaps queries QMK API for a list of keymaps associated with a particular
// keyboard.
func Keymaps(kb string) []string {
	keyboardData := queryQMK(kb)
	return keyboardData[kb].Keymaps
}

// queryQMK is the principal internal function that queries QMK's api
func queryQMK(kb string) map[string]Keyboard {
	var rawJSON json.RawMessage
	var rawData rawData
	var escapedString string

	escapedString = (&url.URL{Path: kb}).String()

	keyboardURL := fmt.Sprintf("%s/%s", qmkAPI, escapedString)

	req, err := http.NewRequest(http.MethodGet, keyboardURL, nil)
	if err != nil {
		log.Print(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("The HTTP request failed with error: %s", err)
	}

	rawJSON, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("The http request could not be read: %s", err)
	}

	err = json.Unmarshal(rawJSON, &rawData)
	if err != nil {
		log.Fatalf("The JSON could not be parsed: %s", err)
	}

	return rawData.Keyboards
}