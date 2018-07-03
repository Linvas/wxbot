package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Linvas/wxbot/wx"
)

func main() {
	uuid, err := wx.GetUUID()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("UUID : " + uuid)
	wx.GenQrFile("./qr321.png", uuid)
	redirectURL := ""
	for {
		redirectURL, err = wx.CheckLogin(uuid)
		if err != nil {
			if strings.Contains(err.Error(), "window.code=400") {
				uuid, err := wx.GetUUID()
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println("UUID : " + uuid)
				wx.GenQrFile("./qr321.png", uuid)
			}
			time.Sleep(time.Second * 3)
			fmt.Println(err.Error())
		} else {
			fmt.Println(redirectURL)
			break
		}
	}
	wx.GetLoginInfo(redirectURL)

}
