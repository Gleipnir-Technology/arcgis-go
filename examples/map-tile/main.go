package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
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

	var lat = flag.Float64("lat", 37.332003, "The latitude to pull the tile from")
	var lng = flag.Float64("lng", -122.0307812, "The longitude to pull the tile from")
	flag.Parse()

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
	for _, s := range map_services {
		log.Info().Str("id", s.ID).Str("name", s.Name).Str("url", s.URL).Str("title", s.Title).Msg("Extracting tiles from map service")
		level := 14
		parent_dir := fmt.Sprintf("tiles/%s", s.Name)
		err := os.MkdirAll(parent_dir, 0750)
		if err != nil {
			log.Error().Err(err).Str("parent_dir", parent_dir).Msg("Failed to make parent dir")
			os.Exit(4)
		}
		for {
			filename := fmt.Sprintf("%s/%d.jpg", parent_dir, level)
			img, err := s.TileGPS(ctx, gis, level, *lat, *lng)
			if err != nil {
				log.Error().Err(err).Msg("tile failure")
				os.Exit(4)
			}
			// Create file in configured directory
			dst, err := os.Create(filename)
			if err != nil {
				log.Error().Err(err).Msg("file create failure")
				os.Exit(4)
			}
			defer dst.Close()
			// Copy rest of request body to file
			_, err = io.Copy(dst, bytes.NewReader(img))
			if err != nil {
				log.Error().Err(err).Msg("file copy failure")
				os.Exit(4)
			}
			log.Info().Str("filename", filename).Msg("Wrote file")
			level += 1
			if level > 30 {
				break
			}
		}
	}
}
