package birdhunter

import (
	"fmt"
	"net/http"
)

type HTTPClient struct {
	Url             string
	Method          string
	ContentType     string
	ContentLanguage string
	Cookies         []*http.Cookie
}

type C struct {
	Name     string
	Password string
	Cookies  []*http.Cookie
}

type Body struct {
	Target Target `json:"tag"`
}

type Target struct {
	Media Media `json:"media"`
	Top   Top   `json:"top_posts"`
}

type Top struct {
	Item []Item `json:"nodes"`
}

type Media struct {
	Item  []Item `json:"nodes"`
	Count int    `json:"count"`
}

type Item struct {
	Id    string `json:"id"`
	Code  string `json:"code"`
	Likes Likes  `json:"likes"`
}

type Likes struct {
	Count int `json:"count"`
}

func (item *Item) Slug() (slug string) {
	return fmt.Sprintf(pic_url, item.Code)
}

func Slug(code string) string {
	return fmt.Sprintf(pic_url, code)
}
