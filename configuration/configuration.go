package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// Properties exposes the configuration value
var Properties *configuration

// configuration is the concrete type of our config.json
type configuration struct {
	ProjectEnv string `json:"projectEnv"`
	AppName    string `json:"appName"`
	AppURI     string `json:"appURI"`
	Version    string `json:"version"`
	Server     struct {
		Address string `json:"address"`
	} `json:"server"`
	Data struct {
		URI    string `json:"uri"`
		DbName string `json:"dbName"`
		DbPath string `json:"dbPath`
	} `json:"data"`
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed Getting Working Directory - Error:"+err.Error())
		return
	}
	err = Update(wd + `/config.json`)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to initialize configuration.Properties - Error:"+err.Error())
	}
}

// Update changes configuration.Properties using values from the filepath
func Update(filepath string) error {
	c := &configuration{}
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return errors.Wrap(err, "Reading the configuration failed - path:"+filepath)
	}

	err = json.Unmarshal(file, c)
	if err != nil {
		return errors.Wrap(err, "failed to update configuration properties - path:"+filepath)
	}
	Properties = c
	return nil
}
