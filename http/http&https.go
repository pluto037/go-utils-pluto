package http

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPClient is a simple HTTP/HTTPS client
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTPClient with a timeout
func NewHTTPClient(timeout time.Duration) *HTTPClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // This allows insecure SSL connections, use it for testing only
	}

	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	return &HTTPClient{client}
}

// Get sends a GET request to the specified URL and returns the response body
func (c *HTTPClient) Get(url string) ([]byte, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post sends a POST request to the specified URL with the given body and returns the response body
func (c *HTTPClient) Post(url string, body []byte) ([]byte, error) {
	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
