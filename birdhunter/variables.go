package birdhunter

var (
	base_url           = "https://www.instagram.com/"
	login_url          = "https://www.instagram.com/accounts/login/?force_classic_login"
	post_ajx_login_url = "https://www.instagram.com/accounts/login/ajax/"
	logout_url         = "https://www.instagram.com/accounts/logout/"
	like_url_tmpl      = "https://www.instagram.com/web/likes/%v/like/"
	unlike_url_tmpl    = "https://www.instagram.com/web/likes/%v/unlike/"
	tag_url            = "https://www.instagram.com/explore/tags/%v/?__a=1"
	pic_url            = "https://www.instagram.com/p/%v"

	//info
	user_agent       = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36"
	accept_language  = "es-ES,es"
	x_www_form       = "application/x-www-form-urlencoded"
	application_json = "application/json"

	c = C{}
)
