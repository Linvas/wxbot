package wx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetLoginInfo(redirectURL string, xcfg *XMLConfig) ([]*http.Cookie, error) {
	resp, err := http.Get(redirectURL + "&fun=new&version=v2")
	if err != nil {
		return nil, fmt.Errorf("访问微信登陆回调URL有误: " + err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	if err := xml.Unmarshal(bodyBytes, xcfg); err != nil {
		return nil, err
	}
	if xcfg.Ret != 0 {
		return nil, fmt.Errorf("Xml结果出错 : " + string(bodyBytes))
	}
	return resp.Cookies(), nil
}
