package httpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	_DEFAULT_TIMEOUT = time.Second * 10
)

var httpClient = &http.Client{
	Timeout: time.Second * 20,
}

func request(ctx context.Context, method, url string, headers map[string]string, body io.Reader, timeout time.Duration) ([]byte, error) {
	if timeout == 0 {
		timeout = _DEFAULT_TIMEOUT
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, body)
	if err != nil {
		return nil, err
	}
	for headerK, headerV := range headers {
		req.Header.Set(headerK, headerV)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("httpCode=%v", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func Get(ctx context.Context, url string, headers map[string]string, params url.Values, timeout time.Duration) ([]byte, error) {
	if p := params.Encode(); p != "" {
		url = url + "?" + p
	}
	return request(ctx, http.MethodGet, url, headers, nil, timeout)
}

// Content-Type : application/x-www-form-urlencoded
func PostForm(ctx context.Context, url string, headers map[string]string, params url.Values, timeout time.Duration) ([]byte, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return request(ctx, http.MethodPost, url, headers, strings.NewReader(params.Encode()), timeout)
}

// Content-Type : application/json;charset=UTF-8
func PostJson(ctx context.Context, url string, headers map[string]string, obj interface{}, timeout time.Duration) ([]byte, error) {
	var body io.Reader
	if obj != nil {
		b, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(b))
	}
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json;charset=UTF-8"
	return request(ctx, http.MethodPost, url, headers, body, timeout)
}
