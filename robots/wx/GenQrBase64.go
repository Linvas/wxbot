package wx

//生成登录二维码图片, 方便在网页上显示

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenQrBase64(uuid string) string {
	content := L_URL + uuid
	img, _ := qr.Encode(content, qr.L, qr.Unicode)
	img, _ = barcode.Scale(img, 300, 300)

	emptyBuff := bytes.NewBuffer(nil) //开辟一个新的空buff缓冲区
	png.Encode(emptyBuff, img)
	dist := make([]byte, 50000)                        //开辟存储空间
	base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
	return "data:image/png;base64," + string(dist)     //输出图片base64(type = []byte)
}
