package wx

import "net/url"

func GetLoginInfo(redirectURL string) {
	u, _ := url.Parse(redirectURL)
	km := u.Query()
	km.Add("fun", "new")
	km.Add("version", "v2")

}
