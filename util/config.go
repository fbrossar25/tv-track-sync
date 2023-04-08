package util

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
)

type TvTrackSyncConfig struct {
	Server struct {
		Port int `yaml:"port" validate:"required,min=1,max=65535"`
	} `yaml:"server"`
	Database struct {
		Host string `yaml:"host" validate:"required,min=1"`
		Port int    `yaml:"port" validate:"required,min=0"`
		Name string `yaml:"dbName" validate:"required,min=1"`
		User string `yaml:"user" validate:"required,min=1"`
		Pwd  string `yaml:"pwd" validate:"required,min=1"`
	} `yaml:"database"`
}

var Config TvTrackSyncConfig

func LoadConfig() {
	f, configFileErr := os.Open("tv-track-sync.yml")
	if configFileErr != nil {
		log.Error().Stack().Err(configFileErr).Msg("Erreur à l'ouverture du fichier de config")
		panic(configFileErr)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	configErr := decoder.Decode(&Config)
	if configErr != nil {
		log.Error().Stack().Err(configErr).Msg("Erreur à la lecture du fichier de config")
		panic(configErr)
	}
	v := validator.New()
	validateConfigErr := v.Struct(Config)
	if validateConfigErr != nil {
		for _, e := range validateConfigErr.(validator.ValidationErrors) {
			log.Error().Stack().Str("configError", fmt.Sprint(e)).Msg("Erreur à la lecture du fichier de config")
		}
		panic(validateConfigErr)
	}
}
