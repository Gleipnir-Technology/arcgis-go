package main

import (
	//"fmt"
	"os"

	//"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	var password = os.Getenv("PASSWORD")
	var username = os.Getenv("USERNAME")

	if password == "" {
		log.Error().Msg("Cannot have empty password")
		os.Exit(1)
	}
	if username == "" {
		log.Error().Msg("Cannot have empty username")
		os.Exit(1)
	}
	log.Info().Str("username", username).Str("password", password).Msg("creds")
}
