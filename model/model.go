package model

type JsCode2SessionResponse struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
}

type WeChatSdkError struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type WechatAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
