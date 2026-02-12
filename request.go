package arcgis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/zerolog"
)

func doGetParamsHeaders(ctx context.Context, client http.Client, host string, path string, params map[string]string, headers map[string]string) ([]byte, error) {
	req_url, err := url.Parse(host + path)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}
	return doGetParamsHeadersFullURL(ctx, client, *req_url, params, headers)
}
func doGetParamsHeadersFullURL(ctx context.Context, client http.Client, req_url url.URL, params map[string]string, headers map[string]string) ([]byte, error) {
	logger := zerolog.Ctx(ctx)
	// Parse the URL

	// Add query parameters if any are provided
	if len(params) > 0 {
		q := req_url.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req_url.RawQuery = q.Encode()
	}

	// Create request with context
	logger.Debug().Str("method", "GET").Str("url", req_url.String()).Msg("Making request")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, req_url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}
	return body, nil
}
func doGetParams(ctx context.Context, client http.Client, host string, path string, params map[string]string) ([]byte, error) {
	return doGetParamsHeaders(ctx, client, host, path, params, map[string]string{})
}
func doGet(ctx context.Context, client http.Client, host string, path string) ([]byte, error) {
	return doGetParams(ctx, client, host, path, map[string]string{})
}
func doGetJSON[T any](ctx context.Context, client http.Client, host string, path string) (*T, error) {
	return doGetJSONParamsHeaders[T](ctx, client, host, path, map[string]string{}, map[string]string{})
}
func doGetJSONParams[T any](ctx context.Context, client http.Client, host string, path string, params map[string]string) (*T, error) {
	return doGetJSONParamsHeaders[T](ctx, client, host, path, params, map[string]string{})
}
func doGetJSONParamsHeaders[T any](ctx context.Context, client http.Client, host string, path string, params map[string]string, headers map[string]string) (*T, error) {
	// Add the 'f=json' param to request a response in json format, if it's not already specified
	_, ok := params["f"]
	if !ok {
		params["f"] = "json"
	}
	// Add accept header, if not present
	_, ok = headers["Accept"]
	if !ok {
		headers["Accept"] = "application/json"
	}
	body, err := doGetParamsHeaders(ctx, client, host, path, params, headers)
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
func reqGet(ctx context.Context, r gisRequestor, path string) ([]byte, error) {
	return reqGetParamsHeaders(ctx, r, path, map[string]string{}, map[string]string{})
}
func reqGetFullURL(ctx context.Context, r gisRequestor, req_url url.URL) ([]byte, error) {
	return reqGetParamsHeadersFullURL(ctx, r, req_url, map[string]string{}, map[string]string{})
}
func reqGetParams(ctx context.Context, r gisRequestor, path string, params map[string]string) ([]byte, error) {
	return reqGetParamsHeaders(ctx, r, path, params, map[string]string{})
}
func reqGetParamsHeaders(ctx context.Context, r gisRequestor, path string, params map[string]string, headers map[string]string) ([]byte, error) {
	req_url, err := url.Parse(r.host + path)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}
	return reqGetParamsHeadersFullURL(ctx, r, *req_url, params, headers)
}
func reqGetParamsHeadersFullURL(ctx context.Context, r gisRequestor, req_url url.URL, params map[string]string, headers map[string]string) ([]byte, error) {
	headers = r.authenticator.addAuthHeaders(ctx, headers)
	return doGetParamsHeadersFullURL(ctx, r.client, req_url, params, headers)
}
func reqGetJSON[T any](ctx context.Context, r gisRequestor, path string) (*T, error) {
	return reqGetJSONParamsHeaders[T](ctx, r, path, map[string]string{}, map[string]string{})
}
func reqGetJSONParams[T any](ctx context.Context, r gisRequestor, path string, params map[string]string) (*T, error) {
	return reqGetJSONParamsHeaders[T](ctx, r, path, params, map[string]string{})
}
func reqGetJSONParamsHeaders[T any](ctx context.Context, r gisRequestor, path string, params map[string]string, headers map[string]string) (*T, error) {
	headers = r.authenticator.addAuthHeaders(ctx, headers)
	return doGetJSONParamsHeaders[T](ctx, r.client, r.host, path, params, headers)
}

func doPostFormParams(ctx context.Context, client http.Client, host string, path string, params map[string]string) (http.Header, []byte, error) {
	return doPostFormParamsHeaders(ctx, client, host, path, params, map[string]string{})
}
func doPostFormParamsHeaders(ctx context.Context, client http.Client, host string, path string, params map[string]string, headers map[string]string) (http.Header, []byte, error) {
	//logger := zerolog.Ctx(ctx)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}
	body := strings.NewReader(data.Encode())

	//logger.Debug().Str("method", "POST").Str("params", mapToString(params)).Str("url", url).Int("status", resp.StatusCode).Msg("Making request")
	return doPostParamsHeaders(ctx, client, host, path, body, headers)
}
func doPostParamsHeaders(ctx context.Context, client http.Client, host string, path string, body io.Reader, headers map[string]string) (http.Header, []byte, error) {
	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, host+path, body)
	if err != nil {
		return nil, nil, fmt.Errorf("creating request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	resp_body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}
	return resp.Header, resp_body, nil
}
func doPostJSON[T any](ctx context.Context, client http.Client, host string, path string, params map[string]string) (*T, error) {
	return doPostJSONParamsHeaders[T](ctx, client, host, path, map[string]string{}, map[string]string{})
}
func doPostFormToJSON[T any](ctx context.Context, client http.Client, host string, path string, params map[string]string, headers map[string]string) (*T, error) {
	//logger := zerolog.Ctx(ctx)
	_, body, err := doPostFormParamsHeaders(ctx, client, host, path, params, headers)
	if err != nil {
		return nil, fmt.Errorf("do POST: %w", err)
	}
	return parseJSON[T](body)
}
func doPostJSONParamsHeaders[T any](ctx context.Context, client http.Client, host string, path string, params map[string]string, headers map[string]string) (*T, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("json marshal: %w", err)
	}
	_, resp_body, err := doPostParamsHeaders(ctx, client, host, path, bytes.NewBuffer(body), headers)
	if err != nil {
		return nil, fmt.Errorf("do post: %w", err)
	}
	return parseJSON[T](resp_body)
}
func parseJSON[T any](body []byte) (*T, error) {
	// Decode JSON
	var result T
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &result, nil
}

func logRequestBase(ctx context.Context, r *http.Request) {
	logger := zerolog.Ctx(ctx)
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
func mapToString(m map[string]string) string {
	var b strings.Builder
	first := true
	for k, v := range m {
		if first {
			b.WriteString(fmt.Sprintf("%s: %s", k, v))
			first = false
		} else {
			b.WriteString(fmt.Sprintf(", %s: %s", k, v))
		}
	}
	return b.String()
}
func reqPostFormToJSON[T any](ctx context.Context, r gisRequestor, path string, params map[string]string) (*T, error) {
	headers := r.authenticator.addAuthHeaders(ctx, map[string]string{})
	return doPostFormToJSON[T](ctx, r.client, r.host, path, params, headers)
}
func reqPostJSONParams[T any](ctx context.Context, r gisRequestor, path string, params map[string]string) (*T, error) {
	headers := r.authenticator.addAuthHeaders(ctx, map[string]string{})
	return doPostJSONParamsHeaders[T](ctx, r.client, r.host, path, params, headers)
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
