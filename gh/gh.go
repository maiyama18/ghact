package gh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type File struct {
	sha     string
	content string
}

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
}

func NewClient() *Client {
	baseURL, _ := url.Parse("https://api.github.com")

	return &Client{
		httpClient: http.DefaultClient,
		baseURL:    baseURL,
	}
}

func (c *Client) Login(username, token string) error {
	reqURL := *c.baseURL
	reqURL.Path = "/users"
	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, token)

	code, err := c.request(req, nil)
	if err != nil {
		return err
	}
	if code < 200 || 300 <= code {
		return fmt.Errorf("ERROR: response status code: %d", code)
	}

	os.Setenv("GHACT_USERNAME", username)
	os.Setenv("GHACT_TOKEN", token)

	return nil
}

func (c *Client) request(req *http.Request, body interface{}) (int, error) {
	resp, err := c.httpClient.Do(req)
	code := resp.StatusCode
	if err != nil {
		return code, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return code, err
	}

	if body != nil {
		if err := json.Unmarshal(b, body); err != nil {
			return code, err
		}
	}

	return code, nil
}
