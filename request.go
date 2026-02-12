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
	logger := zerolog.Ctx(ctx)
	// Parse the URL
	reqURL, err := url.Parse(host + path)
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
	logger.Debug().Str("method", "GET").Str("url", reqURL.String()).Msg("Making request")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
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
	logger := zerolog.Ctx(ctx)
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}
	body := strings.NewReader(data.Encode())
	url := host + path
	resp, err := client.PostForm(url, data)
	logger.Debug().Str("method", "POST").Str("params", mapToString(params)).Str("url", url).Int("status", resp.StatusCode).Msg("Making request")
	if err != nil {
		return nil, nil, fmt.Errorf("make post form: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	resp_body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}
	/*
		for k, v := range resp.Header {
			logger.Debug().Str("header", k).Strs("values", v).Msg("response header")
		}
	*/
	return resp.Header, resp_body, nil
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
	return doPostJSONParams[T](ctx, client, host, path, map[string]string{})
}
func doPostJSONParams[T any](ctx context.Context, client http.Client, host string, path string, params map[string]string) (*T, error) {
	// Set headers
	headers := map[string]string{
		"Accept": "application/json",
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("json marshal: %w", err)
	}
	_, resp_body, err := doPostParamsHeaders(ctx, client, host, path, bytes.NewBuffer(body), headers)
	if err != nil {
		return nil, fmt.Errorf("do post: %w", err)
	}

	// Decode JSON
	var result T
	if err := json.NewDecoder(bytes.NewReader(resp_body)).Decode(&result); err != nil {
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
func reqPostJSONParams[T any](ctx context.Context, r gisRequestor, path string, params map[string]string) (*T, error) {
	return doPostJSONParams[T](ctx, r.client, r.host, path, params)
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
