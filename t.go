package main

import (
	"fmt"

	"github.com/Linvas/wc/wx"
)

func main() {
	resp, err := wx.GetUUID()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
}
