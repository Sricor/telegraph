package telegraph

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type client struct {
	api   string
	token string
}

type Response struct {
	OK     bool            `json:"ok"`
	Error  string          `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

func NewClient() *client {
	return &client{
		api: "https://api.telegra.ph/",
	}
}

func (c *client) SetAPI(api string) {
	c.api = api
}

func (c *client) SetToken(token string) {
	c.token = token
}

func (c *client) Token() (result string) {
	return c.token
}

func (c *client) request(path string, payload interface{}) (response *Response, err error) {
	var (
		body []byte
		resp *http.Response
	)

	body, err = jsonMarshal(payload)
	if err != nil {
		return
	}

	resp, err = http.Post(c.api+path, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	response = &Response{}
	err = jsonUnmarshal(body, response)
	if err != nil {
		return
	}

	if !response.OK {
		err = errors.New(response.Error)
		return
	}

	return
}

// Parameters with token
func (c *client) withToken(v interface{}) (result interface{}) {
	result = &struct {
		AccessToken string `json:"access_token"`
		Parameters  interface{}
	}{c.token, v}
	return
}

func jsonMarshal(value any) ([]byte, error) {
	return json.Marshal(value)
}

func jsonUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
