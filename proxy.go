package arcgis

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/user"
)

func MITMProxyTransport() (*http.Transport, error) {
	// Load mitmproxy certificate
	usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user")
	}
	certPath := usr.HomeDir + "/.mitmproxy/mitmproxy-ca-cert.pem" // Default location
	certBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read mitmproxy cert")
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
	return transport, nil
}
