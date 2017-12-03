package birdhunter

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func client(c HTTPClient, b interface{}) (resp *http.Response, err error) {
	cl := &http.Client{}

	s, _ := json.Marshal(b)
	body := bytes.NewBuffer(s)

	req, err := http.NewRequest(
		c.Method,
		c.Url,
		body,
	)

	if c.Content_type != "" {
		req.Header.Add("Content_type", c.Content_type)
	}

	for _, cookie := range c.Cookies {
		req.AddCookie(cookie)
	}

	return cl.Do(req)
}

func NewClient() {
	res, _ := client(
		HTTPClient{
			Url:    base_url,
			Method: "GET",
		},
		nil,
	)

	c.Cookies = res.Cookies()
}
