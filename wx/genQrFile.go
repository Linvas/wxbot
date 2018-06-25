package wx

//生成登录二维码图片

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func genQrFile(filePath, uuid string) error {
	content := L_URL + uuid
	img, _ := qr.Encode(content, qr.L, qr.Unicode)

	img, _ = barcode.Scale(img, 300, 300)

	return writePng(filePath, img)
}

func writePng(filePath string, img image.Image) error {
	if !createDir(filePath) {
		return fmt.Errorf("创建文件夹失败")
	}
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件失败")
	}
	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("写入文件失败")
	}
	file.Close()
	fmt.Printf("写入文件成功: %s\n", filePath)
	return nil
}

func createDir(filePath string) bool {
	var err error
	// filePath, _ = filepath.Abs(filePath)
	dirPath := filepath.Dir(filePath)
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(dirPath, 0777)
			if err != nil {
				fmt.Println(err.Error())
				return false
			}
		} else {
			fmt.Println(err.Error())
			return false
		}
	} else {
		return dirInfo.IsDir()
	}
	return true
}
