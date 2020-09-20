package minigame

import (
	"fmt"
	"testing"
)

func TestMiniGame_JsCode2Session(t *testing.T) {

	m := &MiniGame{
		AppId:     "wx81425285de3c348b",
		AppSecret: "e82311981cbca8790221323516535334",
	}
	session, e := m.JsCode2Session("073nxJkl2iioE54Zofol2hfxQW2nxJkq")
	fmt.Printf("%#v", session)
	fmt.Printf("%#v", e)
	// {"errcode":40029,"errmsg":"invalid code, hints: [ req_id: KFJeN.iCe-Vh7RGA ]"}
	// {"session_key":"50PDOB89H92Jj99jlylcRQ==","openid":"oDk4r5NH_SUN8JravjpgG3W0wgqM"}

}
