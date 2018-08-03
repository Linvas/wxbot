package wx

import (
	"fmt"
	"net/url"
	"time"
)

func WebInit(xcfg *XMLConfig) {
	var timestamp = time.Now().UnixNano() / 1000000
	params := url.Values{}
	params.Add("pass_ticket", xcfg.PassTicket)
	params.Add("skey", xcfg.Skey)
	params.Add("r", fmt.Sprintf("%d", timestamp))
}
