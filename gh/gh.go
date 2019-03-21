package gh

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type File struct {
	Sha     string `json:"sha"`
	Content string `json:"content"`
}

func (f *File) Decode() error {
	b, err := base64.StdEncoding.DecodeString(f.Content)
	if err != nil {
		return err
	}

	f.Content = string(b)
	return nil
}

type Commit struct {
	Sha     string `json:"sha"`
	Content string `json:"content"`
	Message string `json:"message"`
}

func (c *Commit) Encode() {
	c.Content = base64.StdEncoding.EncodeToString([]byte(c.Content))
}

func NewCommit(sha, content, message string) *Commit {
	return &Commit{sha, content, message}
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

	return c.request(req, nil)
}

func (c *Client) Fetch(owner, repo, filePath string) (*File, error) {
	reqURL := *c.baseURL
	reqURL.Path = "/" + path.Join("repos", owner, repo, "contents", filePath)
	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var file File
	err = c.request(req, &file)
	if err != nil {
		return nil, err
	}
	if err := file.Decode(); err != nil {
		return nil, err
	}

	return &file, nil
}

func (c *Client) Create(owner, repo, filePath, token string, commit *Commit) error {
	commit.Encode()
	b, err := json.Marshal(commit)

	reqURL := *c.baseURL
	reqURL.Path = "/" + path.Join("repos", owner, repo, "contents", filePath)
	req, err := http.NewRequest(http.MethodPut, reqURL.String(), bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+token)

	return c.request(req, nil)
}

func (c *Client) Update(owner, repo, filePath, token string, commit *Commit) error {
	fmt.Println("UPDATE")
	commit.Encode()
	b, err := json.Marshal(commit)
	fmt.Println(b)

	reqURL := *c.baseURL
	reqURL.Path = "/" + path.Join("repos", owner, repo, "contents", filePath)
	req, err := http.NewRequest(http.MethodPut, reqURL.String(), bytes.NewBuffer(b))
	fmt.Println(req)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+token)

	return c.request(req, nil)
}

func (c *Client) request(req *http.Request, body interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	code := resp.StatusCode
	if code < 200 || 300 <= code {
		return fmt.Errorf("ERROR: response status code: %d", code)
	}

	if body != nil {
		if err := json.Unmarshal(b, body); err != nil {
			return err
		}
	}

	return nil
}
