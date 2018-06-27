package wx

const (
	BASE_URL   = "https://login.weixin.qq.com"           /* API基准地址 */
	UUID_URL   = BASE_URL + "/jslogin?"                  /* 获取uuid的地址 */
	QRCODE_URL = BASE_URL + "/qrcode/"                   /* 获取二维码的地址 */
	L_URL      = BASE_URL + "/l/"                        /* 二维码地址内容 */
	LOGIN_URL  = BASE_URL + "/cgi-bin/mmwebwx-bin/login?" /* 登陆URL  */

	API_BASE_URL              = "https://wx.qq.com/cgi-bin/mmwebwx-bin" /* API基准地址 */
	INIT_URL                  = API_BASE_URL + "/webwxinit"             /* 初始化URL  */
	STATUS_NOTIFY_URL         = API_BASE_URL + "/webwxstatusnotify"     /* 通知微信状态变化 */
	GET_ALL_CONTACT_URL       = API_BASE_URL + "/webwxgetcontact"       /* 获取所有联系人信息 */
	WEB_WX_SYNC_URL           = API_BASE_URL + "/webwxsync"             /* 拉取同步消息 */
	WEB_WX_SENDMSG_URL        = API_BASE_URL + "/webwxsendmsg"          /* 发送消息 */
	WEB_WX_UPDATECHATROOM_URL = API_BASE_URL + "/webwxupdatechatroom"   /* 群更新，主要用来邀请好友 */

	SYNC_CHECK_URL = "https://webpush.wx.qq.com/cgi-bin/mmwebwx-bin/synccheck" /* 检查心跳URL */
)

const (
	JSON_HEADER = "application/json;charset=UTF-8"
)

const (
	APPID_WEB    = "wx782c26e4c19acffb"
	APPID_PC     = ""
	APPID_MACOSX = ""
	LANG         = "zh_CN"
	UA_WEB       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3419.0 Safari/537.36"
	UA_ANDROID   = "" //TODO
	UA_IOS       = "" //TODO
	UA_PC        = "" //TODO PC电脑客户端UA
)

type XmlConfig struct {
}
