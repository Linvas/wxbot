package wx

import (
	"fmt"
	"time"
)

//GetRedirectURL 从扫描二维码从微信获取重定向地址
func GetRedirectURL(uuid string) (string, error) {
	fmt.Println("正在验证登陆... ...")
	redirectURL := ""
	var err error
	for {
		redirectURL, err = CheckLogin(uuid)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}
	fmt.Println(redirectURL)
	return redirectURL, nil
}
