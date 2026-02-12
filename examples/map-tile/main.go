package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	var base = os.Getenv("BASEURL")
	var password = os.Getenv("PASSWORD")
	var username = os.Getenv("USERNAME")

	if base == "" {
		base = "https://www.arcgis.com"
	}
	if password == "" {
		log.Error().Msg("Cannot have empty password")
		os.Exit(1)
	}
	if username == "" {
		log.Error().Msg("Cannot have empty username")
		os.Exit(1)
	}
	ctx := context.TODO()
	ctx = log.With().Str("component", "arcgis").Logger().WithContext(ctx)

	// Load mitmproxy certificate
	usr, err := user.Current()
	if err != nil {
		log.Error().Err(err).Msg("current usr")
		os.Exit(1)
	}
	certPath := usr.HomeDir + "/.mitmproxy/mitmproxy-ca-cert.pem" // Default location
	certBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		panic(err)
	}

	// Create a certificate pool and add the certificate
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	rootCAs.AppendCertsFromPEM(certBytes)

	// Configure proxy
	proxyURL, _ := url.Parse("http://127.0.0.1:8080")

	// Configure transport with proxy and certificates
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{
			RootCAs: rootCAs,
		},
	}

	gis, err := arcgis.NewArcGISTransport(ctx, &base, &arcgis.AuthenticatorUsernamePassword{
		Password: password,
		Username: username,
	}, transport)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create arcgis")
		os.Exit(2)
	}
	map_services, err := gis.MapServices(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get map service")
		os.Exit(3)
	}
	log.Info().Int("total", map_services.Total).Msg("got results")
	for _, s := range map_services.Results {
		log.Info().Str("id", s.ID).Str("name", s.Name).Str("url", s.URL).Msg("Map service")
	}
	log.Info().Str("username", username).Str("password", password).Msg("creds")
}
