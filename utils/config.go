package utils

import (
	"bytes"
	"log"
	"strings"

	"github.com/fitzix/assassin/models"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var appConf models.Config

func InitConf() {
	v := viper.New()
	v.SetConfigType("yaml")
	b, err := yaml.Marshal(models.NewConfig())
	if err != nil {
		log.Fatalf("maeshal default err: %s", err)
	}
	if err := v.MergeConfig(bytes.NewReader(b)); err != nil {
		log.Fatalf("merge default err: %s", err)
	}
	v.SetConfigFile("config.yml")
	if err := v.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			log.Fatalf("merge default err: %s", err)
		}
		// dont return error if file is missing. overwrite file is optional
	}
	// tell viper to overwrite env variables
	v.AutomaticEnv()
	v.SetEnvPrefix("ASN")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.Unmarshal(&appConf); err != nil {
		log.Fatalf("unmarshal default err: %s", err)
	}
}


func GetConf() models.Config {
	return appConf
}
