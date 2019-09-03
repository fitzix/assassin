package utils

import (
	"bytes"
	"strings"

	"github.com/fitzix/assassin/models"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func LoadConfig(config *models.Config) error {
	v := viper.New()
	v.SetConfigType("yaml")
	b, err := yaml.Marshal(models.NewConfig())
	if err != nil {
		return err
	}
	if err := v.MergeConfig(bytes.NewReader(b)); err != nil {
		return err
	}
	v.SetConfigFile("config.yml")
	if err := v.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return err
		}
		// dont return error if file is missing. overwrite file is optional
	}
	// tell viper to overwrite env variables
	v.AutomaticEnv()
	v.SetEnvPrefix("ASN")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return v.Unmarshal(config)
}
