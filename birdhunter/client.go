package birdhunter

import (
	"bytes"
	"encoding/json"
	"io"
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

	if c.ContentType != "" {
		req.Header.Add("Content-Type", c.ContentType)
	}

	if c.ContentLanguage != "" {
		req.Header.Add("Content-Language", c.ContentLanguage)
	}

	for _, cookie := range c.Cookies {
		req.AddCookie(cookie)
	}

	return cl.Do(req)
}

func getJson(body io.ReadCloser, target interface{}) error {
	defer body.Close()

	return json.NewDecoder(body).Decode(target)
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
