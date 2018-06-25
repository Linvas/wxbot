package wx

/* 判断微信是否登陆 */
func CheckLogin(uuid string) (int64, string) {
	var timestamp int64 = time.Now().UnixNano() / 1000000
	paraMap := e.GetLoginParaEnum()
	paraMap[e.UUID] = uuid
	paraMap[e.TimeStamp] = fmt.Sprintf("%d", timestamp)
	paraMap[e.R] = fmt.Sprintf("%d", ^(int32)(timestamp))
	var paraStr = t.GetURLParams(paraMap)

	resp, err := http.Get(e.LOGIN_URL + paraStr)
	if err != nil {
		return 0, "访问微信服务器API失败:" + err.Error()
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "获取微信API反馈登陆数据失败:" + err.Error()
	}

	var resultStr = string(bodyBytes)

	reg := regexp.MustCompile(`^window.code=(\d+);`)
	matches := reg.FindStringSubmatch(resultStr)
	if len(matches) < 2 {
		return 0, "预期的返回结果格式不匹配"
	}

	status, err := strconv.ParseInt(matches[1], 10, 64)

	return status, resultStr
}

// Login: login api
func (api *ApiV2) Login(common *Common, uuid, tip string) (string, error) {
	km := url.Values{}
	km.Add("tip", tip)
	km.Add("uuid", uuid)
	km.Add("r", strconv.FormatInt(time.Now().Unix(), 10))
	km.Add("_", strconv.FormatInt(time.Now().Unix(), 10))
	uri := common.LoginUrl + "/cgi-bin/mmwebwx-bin/login?" + km.Encode()
	body, _ := api.httpClient.Get(uri, nil)
	strb := string(body)
	if strings.Contains(strb, "window.code=200") &&
		strings.Contains(strb, "window.redirect_uri") {
		ss := strings.Split(strb, "\"")
		if len(ss) < 2 {
			return "", fmt.Errorf("parse redirect_uri fail, %s", strb)
		}
		return ss[1], nil
	}

	return "", fmt.Errorf("login response, %s", strb)
}