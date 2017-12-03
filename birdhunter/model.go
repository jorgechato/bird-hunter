package birdhunter

import "net/http"

type HTTPClient struct {
	Url          string
	Method       string
	Content_type string
	Cookies      []*http.Cookie
}

type Client struct {
	Name     string
	Password string
	Cookies  []*http.Cookie
}
