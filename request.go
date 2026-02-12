package arcgis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type gisRequestor struct {
	authenticator Authenticator
	client        http.Client
	host          string
}

func newGisRequestor(authenticator Authenticator, host string) gisRequestor {
	return gisRequestor{
		authenticator: authenticator,
		client:        http.Client{},
		host:          host,
	}

}
func (r gisRequestor) withHost(host string) gisRequestor {
	return gisRequestor{
		authenticator: r.authenticator,
		client:        http.Client{},
		host:          host,
	}
}

func doGetParamsHeaders(ctx context.Context, r gisRequestor, path string, params map[string]string, headers map[string]string) ([]byte, error) {
	// Parse the URL
	reqURL, err := url.Parse(r.host + path)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	// Add query parameters if any are provided
	if len(params) > 0 {
		q := reqURL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		reqURL.RawQuery = q.Encode()
	}

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// Make the request
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}
	return body, nil
}
func doGetParams(ctx context.Context, r gisRequestor, path string, params map[string]string) ([]byte, error) {
	return doGetParamsHeaders(ctx, r, path, params, map[string]string{})
}
func doGet(ctx context.Context, r gisRequestor, path string) ([]byte, error) {
	return doGetParams(ctx, r, path, map[string]string{})
}
func doJSONGet[T any](ctx context.Context, r gisRequestor, path string) (*T, error) {
	return doJSONGetParams[T](ctx, r, path, map[string]string{})
}
func doJSONGetParams[T any](ctx context.Context, r gisRequestor, path string, params map[string]string) (*T, error) {
	body, err := doGetParamsHeaders(ctx, r, path, params, map[string]string{"Accept": "application/json"})
	if err != nil {
		return nil, fmt.Errorf("doing request: %w", err)
	}
	// During normal operation the ArcGIS server does _not_ respond with the standard
	// 400-level response codes on error. Instead it responds with 200, which is wrong,
	// but we can't force them to be standards-compliant. Instead, we have to attempt
	// to parse an error. If it works we have an error. If not, it must be something else.
	errorFromJSON := tryParseError(ctx, body)
	if errorFromJSON != nil {
		return nil, fmt.Errorf("response was an application-level error in JSON: %w", errorFromJSON)
	}

	// Decode JSON
	var result T
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &result, nil
}

func doJSONPost[T any](ctx context.Context, r gisRequestor, path string, params map[string]string) (*T, error) {

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.host+path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")

	// Make the request
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	// Decode JSON
	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &result, nil
}

func logRequestBase(ctx context.Context, r *http.Request) {
	logger := LoggerFromContext(ctx)
	if r == nil {
		logger.Warn().Msg("Can't log request, it's nil")
		return
	}

	// Create a copy of the URL to avoid modifying the original
	cleanURL := *r.URL

	// Remove query parameters
	//cleanURL.RawQuery = ""
	q, _ := url.ParseQuery(cleanURL.RawQuery)
	q.Del("token")
	cleanURL.RawQuery = q.Encode()

	logger.Debug().Str("method", r.Method).Str("url", cleanURL.String()).Msg("ArcGIS request")
}

func tryParseError(ctx context.Context, data []byte) error {
	var msg ErrorResponse
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON: %w", err)
	}
	if msg.Error.Code != 0 || msg.Error.Message != "" || len(msg.Error.Details) > 0 {
		return newAPIError(ctx, msg)
	}
	return nil
}
