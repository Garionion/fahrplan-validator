package main

import (
	"github.com/Garionion/fahrplan"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/url"
)

type Configuration struct {
	Path string `yaml:"Path" env:"SCHEDULE" env-default:"./schedule.json"`
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func main() {
	var cfg Configuration
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Printf("Can't read configuration: %v\n", err)
	}
	if isValidUrl(cfg.Path) {
		_, err := fahrplan.GetScheduleFromWeb(cfg.Path)
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		_, err := fahrplan.GetScheduleFromFile(cfg.Path)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	log.Printf("Everything seems to be fine. Parsing of Schedule was successfull")
}
