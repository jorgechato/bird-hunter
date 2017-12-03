package birdhunter

func Login(user string, pass string) {
	c.Name = user
	c.Password = pass

	client(
		HTTPClient{
			Url:          login_url,
			Method:       "POST",
			Content_type: x_www_form,
			Cookies:      c.Cookies,
		},
		c,
	)
}

//TODO: get list of items by hashtag
