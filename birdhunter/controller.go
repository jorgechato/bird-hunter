package birdhunter

import "fmt"

func Login(user string, pass string) {
	c.Name = user
	c.Password = pass

	Client(
		HTTPClient{
			Url:         login_url,
			Method:      "POST",
			ContentType: x_www_form,
			Cookies:     c.Cookies,
		},
		c,
	)
}

func GetItemsByTag(tag string) (Target, error) {
	url := fmt.Sprintf(tag_url, tag)

	res, err := Client(
		HTTPClient{
			Url:             url,
			Method:          "GET",
			ContentType:     application_json,
			ContentLanguage: accept_language,
		},
		nil,
	)

	body := Body{}
	err = getJson(res.Body, &body)

	return body.Target, err
}
