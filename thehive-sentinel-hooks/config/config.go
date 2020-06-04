package config

import (
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// Conf struct is used to parse the config.yml file
type Conf struct {
	ResolvedCaseURL  string `yaml:"resolvedCaseURL"`
	IgnoredAlertURL  string `yaml:"ignoredAlertURL"`
	ImportedAlertURL string `yaml:"importedAlertURL"`
	NewCaseURL       string `yaml:"newCaseURL"`
	NewAlertURL      string `yaml:"newAlertURL"`
	Organization     string `yaml:"organization"`
	LogLevel         string `yaml:"logLevel"`
}

// GetConfig parses config.yml configuration file
func GetConfig(c *Conf, f *string) error {
	yamlFile, err := ioutil.ReadFile(*f)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		return err
	}
	log.Info().Msgf("config.yml imported from: %s", *f)
	return nil
}
