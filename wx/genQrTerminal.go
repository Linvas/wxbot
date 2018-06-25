package wx

//在终端显示二维码

import (
	"os"

	"github.com/mdp/qrterminal"
)

func GenQrTerminal(uuid string) {
	qrterminal.Generate(L_URL+uuid, qrterminal.L, os.Stdout)
}
