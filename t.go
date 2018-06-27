package main

import (
	"fmt"

	"github.com/Linvas/wxbot/wx"
)

func main() {
	uuid, err := wx.GetUUID()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("UUID : " + uuid)
	wx.GenQrFile("./qr321.png", uuid)

	redirectURL, err := wx.CheckLogin(uuid)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(redirectURL)
	}
}
