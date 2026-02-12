package arcgis

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
)

type gisRequestor struct {
	authenticator Authenticator
	client        http.Client
	host          string
}

func newGisRequestor(ctx context.Context, authenticator Authenticator, host string, transport *http.Transport) (gisRequestor, error) {
	logger := zerolog.Ctx(ctx)
	if transport == nil {
		transport = &http.Transport{}
	}
	client := http.Client{
		Transport: transport,
	}
	err := authenticator.init(ctx, client)
	if err != nil {
		return gisRequestor{}, fmt.Errorf("init auth: %w", err)
	}
	f_headers := authenticator.addAuthHeaders(ctx, map[string]string{})
	logger.Debug().Str("access_token", mapToString(f_headers)).Msg("handing auth to requestor")
	return gisRequestor{
		authenticator: authenticator,
		client:        client,
		host:          host,
	}, nil
}
func (r gisRequestor) withHost(host string) gisRequestor {
	return gisRequestor{
		authenticator: r.authenticator,
		client:        http.Client{},
		host:          host,
	}
}

func (r gisRequestor) doGetParamsHeaders(ctx context.Context, path string, params map[string]string, headers map[string]string) ([]byte, error) {
	headers = r.authenticator.addAuthHeaders(ctx, headers)
	return doGetParamsHeaders(ctx, r.client, r.host, path, params, headers)
}
func (r gisRequestor) doGetParams(ctx context.Context, path string, params map[string]string) ([]byte, error) {
	return doGetParamsHeaders(ctx, r.client, r.host, path, params, map[string]string{})
}
func (r gisRequestor) doGet(ctx context.Context, path string) ([]byte, error) {
	return doGetParamsHeaders(ctx, r.client, r.host, path, map[string]string{}, map[string]string{})
}
