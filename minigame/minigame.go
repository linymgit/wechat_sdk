package minigame

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/linymgit/wechat_sdk/model"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// wx81425285de3c348b
// e82311981cbca8790221323516535334
type MiniGame struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	lock      sync.Mutex
	model.WechatAccessToken
	AccessTokenTime time.Time
}

func NewMiniGame(appId, appSecret string) *MiniGame {
	return &MiniGame{
		AppId:     appId,
		AppSecret: appSecret,
		lock:      sync.Mutex{},
	}
}

// GET https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
func (m *MiniGame) GetAccessToken() (r *model.WechatAccessToken, err error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	if m.AccessToken != "" && time.Now().Before(m.AccessTokenTime.Add(time.Duration(m.ExpiresIn-300)*time.Second)) {
		return &m.WechatAccessToken, nil
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", m.AppId, m.AppSecret)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return
	}
	if r.AccessToken == "" {
		w := model.WeChatSdkError{}
		err = json.Unmarshal(bytes, &w)
		if err != nil {
			return
		}
		err = errors.New(w.Errmsg)
		return
	}
	m.WechatAccessToken = *r
	m.AccessTokenTime = time.Now()
	return

}

// GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
func (m *MiniGame) JsCode2Session(code string) (r *model.JsCode2SessionResponse, err error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		m.AppId, m.AppSecret, code)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return
	}
	if r.Openid == "" {
		w := model.WeChatSdkError{}
		err = json.Unmarshal(bytes, &w)
		if err != nil {
			return
		}
		err = errors.New(w.Errmsg)
		return
	}
	return

}
