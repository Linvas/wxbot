package wx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//GetUUID 从微信获取UUID
func GetUUID() (string, error) {
	km := url.Values{}
	km.Add("appid", APPID_WEB)
	km.Add("fun", "new")
	km.Add("lang", LANG)
	km.Add("_", fmt.Sprintf("%v", time.Now().Unix()))

	resp, err := http.Get(UUID_URL + km.Encode())
	fmt.Println(UUID_URL + km.Encode())
	if err != nil {
		//TODO 此处应该打印日志
		return "", fmt.Errorf("访问微信服务器API失败: %s", err.Error())
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("获取微信API反馈UUID数据失败: %s", err.Error())
	}

	ss := strings.Split(string(bodyBytes), "\"")
	if len(ss) < 2 {
		return "", fmt.Errorf("获取UUID信息无效, %s", string(bodyBytes))
	}
	resp.Body.Close()
	return ss[1], nil
}
