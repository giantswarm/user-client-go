package client

import (
	"github.com/catalyst-zero/api-schema"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ENDPOINT_URL = "%s/%s%s" // {endpoint}/{version}%s
)

// Dial returns a client for the given server, using the given version.
func Dial(server, version string) (client *Client, err error) {
	client = new(Client)

	client.version = version

	client.endpoint, err = parseEndpoint(server)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type Client struct {
	version  string
	endpoint *Endpoint

	LogGetRequest    func(url string, resp *http.Response, err error)
	LogPostRequest   func(url, contentType string, resp *http.Response, err error)
	LogDeleteRequest func(url string, resp *http.Response, err error)
}

func (c *Client) endpointUrl(url string) string {
	return fmt.Sprintf(ENDPOINT_URL, c.endpoint.String(), c.version, url)
}

// NOTE: These need to be redone, since one accepts a result object, the other does not and so on. Very inconsequent. :(
func (c *Client) get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if c.LogGetRequest != nil {
		c.LogGetRequest(url, resp, err)
	}
	return resp, Mask(err)
}

func (c *Client) post(url, contentType string, body io.Reader) (*http.Response, error) {
	resp, err := http.Post(url, contentType, body)
	if c.LogPostRequest != nil {
		c.LogPostRequest(url, contentType, resp, err)
	}
	return resp, Mask(err)
}

// postJson transforms the body into a JSON stream and sends it to the given URL as a HTTP POST request.
func (c *Client) postJSON(url string, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, Mask(err)
	}

	resp, err := c.post(url, "application/json", bytes.NewReader(data))
	return resp, Mask(err)
}

func (c *Client) delete(url string) (*http.Response, error) {
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, Mask(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if c.LogDeleteRequest != nil {
		c.LogDeleteRequest(url, resp, err)
	}

	return resp, Mask(err)
}

func (c *Client) postSchemaJSON(urlFragment string, payload interface{}) (*apischema.Response, error) {
	resp, err := apischema.FromHTTPResponse(c.postJSON(c.endpointUrl(urlFragment), payload))
	return resp, Mask(err)
}

func (c *Client) postSchema(urlFragment string) (*apischema.Response, error) {
	resp, err := apischema.FromHTTPResponse(c.post(c.endpointUrl(urlFragment), "", nil))
	return resp, Mask(err)
}
