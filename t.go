package main

import (
	"fmt"

	"github.com/Linvas/wxbot/wx"
)

func main() {
	resp, err := wx.GetUUID()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)

	wx.CheckLogin()
}
