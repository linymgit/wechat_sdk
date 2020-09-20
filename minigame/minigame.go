package minigame

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/linymgit/wechat_sdk/model"
	"io/ioutil"
	"net/http"
)

// wx81425285de3c348b
// e82311981cbca8790221323516535334
type MiniGame struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
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
