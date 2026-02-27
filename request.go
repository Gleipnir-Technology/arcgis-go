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

/* Raw GET functions */
func doGet(ctx context.Context, client http.Client, host string, path string) ([]byte, error) {
	return doGetParams(ctx, client, host, path, map[string]string{})
}
func doGetParams(ctx context.Context, client http.Client, host string, path string, params map[string]string) ([]byte, error) {
	return doGetParamsHeaders(ctx, client, host, path, params, map[string]string{})
}
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

/* GET JSON interpreter functions */
func doGetJSON[T any](ctx context.Context, client http.Client, host string, path string) (*T, error) {
	return doGetJSONParams[T](ctx, client, host, path, map[string]string{})
}
func doGetJSONParams[T any](ctx context.Context, client http.Client, host string, path string, params map[string]string) (*T, error) {
	return doGetJSONParamsHeaders[T](ctx, client, host, path, params, map[string]string{})
}
func doGetJSONParamsHeaders[T any](ctx context.Context, client http.Client, host string, path string, params map[string]string, headers map[string]string) (*T, error) {
	full_url, err := url.Parse(host + path)
	if err != nil {
		return nil, fmt.Errorf("failed url parse: %w", err)
	}
	return doGetJSONParamsHeadersFullURL[T](ctx, client, *full_url, params, headers)
}
func doGetJSONParamsHeadersFullURL[T any](ctx context.Context, client http.Client, req_url url.URL, params map[string]string, headers map[string]string) (*T, error) {

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
	body, err := doGetParamsHeadersFullURL(ctx, client, req_url, params, headers)
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

/* GET functions using a requestor */
func reqGet(ctx context.Context, r gisRequestor, path string) ([]byte, error) {
	return reqGetParamsHeaders(ctx, r, path, map[string]string{}, map[string]string{})
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

/* GET functions that interpret from JSON */
func reqGetJSON[T any](ctx context.Context, r gisRequestor, path string) (*T, error) {
	return reqGetJSONParamsHeaders[T](ctx, r, path, map[string]string{}, map[string]string{})
}
func reqGetJSONFullURL[T any](ctx context.Context, r gisRequestor, req_url url.URL) (*T, error) {
	return reqGetJSONParamsHeadersFullURL[T](ctx, r, req_url, map[string]string{}, map[string]string{})
}
func reqGetJSONParams[T any](ctx context.Context, r gisRequestor, path string, params map[string]string) (*T, error) {
	return reqGetJSONParamsHeaders[T](ctx, r, path, params, map[string]string{})
}
func reqGetJSONParamsHeaders[T any](ctx context.Context, r gisRequestor, path string, params map[string]string, headers map[string]string) (*T, error) {
	req_url, err := url.Parse(r.host + path)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}
	return reqGetJSONParamsHeadersFullURL[T](ctx, r, *req_url, params, headers)
}
func reqGetJSONParamsHeadersFullURL[T any](ctx context.Context, r gisRequestor, req_url url.URL, params map[string]string, headers map[string]string) (*T, error) {
	headers = r.authenticator.addAuthHeaders(ctx, headers)
	return doGetJSONParamsHeadersFullURL[T](ctx, r.client, req_url, params, headers)
}

/* Raw POST functions */
func doPostRaw(ctx context.Context, client http.Client, req_url url.URL, body io.Reader, headers map[string]string) (http.Header, []byte, error) {
	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, req_url.String(), body)
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

/* POST with form-encoded data */
func doPostForm(ctx context.Context, client http.Client, host string, path string) (http.Header, []byte, error) {
	return doPostFormParams(ctx, client, host, path, map[string]string{})
}
func doPostFormParams(ctx context.Context, client http.Client, host string, path string, params map[string]string) (http.Header, []byte, error) {
	return doPostFormParamsHeaders(ctx, client, host, path, params, map[string]string{})
}
func doPostFormParamsHeaders(ctx context.Context, client http.Client, host string, path string, params map[string]string, headers map[string]string) (http.Header, []byte, error) {
	req_url, err := url.Parse(host + path)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing URL: %w", err)
	}
	return doPostFormParamsHeadersFullURL(ctx, client, *req_url, params, headers)
}
func doPostFormParamsHeadersFullURL(ctx context.Context, client http.Client, req_url url.URL, params map[string]string, headers map[string]string) (http.Header, []byte, error) {
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}
	body := strings.NewReader(data.Encode())
	return doPostRaw(ctx, client, req_url, body, headers)
}

/* POST with form-encoded data, convert from JSON */
func doPostJSON[T any](ctx context.Context, client http.Client, host string, path string) (*T, error) {
	return doPostJSONParamsHeaders[T](ctx, client, host, path, map[string]string{}, map[string]string{})
}
func doPostJSONParamsHeaders[T any](ctx context.Context, client http.Client, host string, path string, params map[string]string, headers map[string]string) (*T, error) {
	req_url, err := url.Parse(host + path)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}
	return doPostJSONParamsHeadersFullURL[T](ctx, client, *req_url, params, headers)
}
func doPostJSONParamsHeadersFullURL[T any](ctx context.Context, client http.Client, req_url url.URL, params map[string]string, headers map[string]string) (*T, error) {
	//logger := zerolog.Ctx(ctx)
	_, ok := params["f"]
	if !ok {
		params["f"] = "json"
	}
	_, body, err := doPostFormParamsHeadersFullURL(ctx, client, req_url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("do POST: %w", err)
	}
	return parseJSON[T](body)
}
func parseJSON[T any](body []byte) (*T, error) {
	// Decode JSON
	var result T
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&result); err != nil {
		return nil, fmt.Errorf("parseJSON: %w", err)
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

/* POST with form-encoded data, using requestor (with auth), convert from JSON */
func reqPostFormToJSON[T any](ctx context.Context, r gisRequestor, path string, params map[string]string) (*T, error) {
	req_url, err := url.Parse(r.host + path)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}
	return reqPostFormToJSONFullURL[T](ctx, r, *req_url, params)
}
func reqPostFormToJSONFullURL[T any](ctx context.Context, r gisRequestor, req_url url.URL, params map[string]string) (*T, error) {
	headers := r.authenticator.addAuthHeaders(ctx, map[string]string{})
	return doPostJSONParamsHeadersFullURL[T](ctx, r.client, req_url, params, headers)
}

/* Utilities */
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

/* New Resty stuff starts here */
