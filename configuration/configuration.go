package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
)

// Config is the concrete type of our config.json
type Config struct {
	ProjectEnv string `json:"projectEnv"`
	AppName    string `json:"appName"`
	AppURI     string `json:"appURI"`
	Version    string `json:"version"`
	Server     struct {
		Address string `json:"address"`
	} `json:"server"`
	Database struct {
		URI  string `json:"uri"`
		Name string `json:"name"`
	} `json:"database"`
	Logger struct {
		Info  *log.Logger
		Warn  *log.Logger
		Error *log.Logger
	} `json:"-"`
}

func FromFile(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "Reading the configuration failed - path:"+path)
	}

	c := &Config{}
	err = json.Unmarshal(file, c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update configuration properties - path:"+path)
	}
	c.Logger.Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	c.Logger.Warn = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	c.Logger.Error = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	return c, nil
}
