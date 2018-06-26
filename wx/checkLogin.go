package wx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//CheckLogin 检查微信是否已经登录
func CheckLogin(uuid string) (string, error) {
	var timestamp = time.Now().UnixNano() / 1000000
	km := url.Values{}
	km.Add("tip", "0")
	km.Add("uuid", uuid)
	km.Add("loginicon", "true")
	km.Add("r", fmt.Sprintf("%d", timestamp))
	km.Add("_", fmt.Sprintf("%d", ^(int32)(timestamp)))
	resp, err := http.Get(LOGIN_URL + km.Encode())
	if err != nil {
		return "", fmt.Errorf("访问微信服务器API失败: " + err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("获取微信API反馈登陆数据失败: " + err.Error())
	}
	body := string(bodyBytes)
	if strings.Contains(body, "window.code=200") &&
		strings.Contains(body, "window.redirect_uri") {
		ss := strings.Split(body, "\"")
		if len(ss) < 2 {
			return "", fmt.Errorf("解析redirect_uri 失败, %s", body)
		}
		return ss[1], nil
	}

	return "", fmt.Errorf("登录响应: %s", body)
}
