package utils

import (
	"net"
	"crypto/md5"
	"fmt"
)

//将本机所有mac地址转换成md5加密
func MacAddrMd5() (string, error) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	macAddr := ""
	for _, inter := range interfaces {
		//fmt.Println(inter.Name)
		if inter.Flags&net.FlagUp != 0 {
			mac := inter.HardwareAddr.String() //获取本机MAC地址
			if mac != "" {
				macAddr += mac
			}
		}
	}
	return encodeMacAddr(macAddr), nil
}

func encodeMacAddr(macAddr string) string {
	has := md5.Sum([]byte(macAddr))

	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
