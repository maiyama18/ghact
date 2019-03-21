package gh

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
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

func (c *Client) Request(req *http.Request, body interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, body); err != nil {
		return err
	}

	return nil
}

//func (c *Client) GetContents(filepath string) (*File, error) {
//	url := *c.baseURL
//	url.Path = "/repos/"
//}
