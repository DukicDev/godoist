package todoist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://api.todoist.com/rest/v2/"

type Client struct {
	APIToken string
}

func NewClient(apiToken string) *Client {
	return &Client{
		APIToken: apiToken,
	}
}

func (client *Client) makeRequest(method, endpoint string, query map[string]string, body interface{}) ([]byte, error) {
	url := baseUrl + endpoint
	var reqBody []byte
	var err error

	if method != "GET" && body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+client.APIToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API returned error: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
