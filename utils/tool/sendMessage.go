package tool

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"qiublog/utils"
)

type body struct {
	Agentid string `json:"agentid"`
	Touser  string `json:"touser"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string `json:"msgtype"`
}
type res struct {
	Errcode        int    `json:"errcode"`
	Errmsg         string `json:"errmsg"`
	Invaliduser    string `json:"invaliduser"`
	Invalidparty   string `json:"invalidparty"`
	Invalidtag     string `json:"invalidtag"`
	Unlicenseduser string `json:"unlicenseduser"`
	Msgid          string `json:"msgid"`
	ResponseCode   string `json:"response_code"`
}

type tooken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func WxPush(content string) error {
	data := body{
		Agentid: utils.Config.Server.Push.WxPushAgentid,
		Touser:  "@all",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
		Msgtype: "text",
	}
	tookenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", utils.Config.Server.Push.WxPushCorpID, utils.Config.Server.Push.WxPushSecret)
	get, err := http.Get(tookenUrl)
	if err != nil {
		return errors.New("WxPush:Token Get error")
	}
	defer get.Body.Close()
	tk, err := io.ReadAll(get.Body)
	if err != nil {
		return errors.New("WxPush:Token io.ReadAll error")
	}
	Tooken := &tooken{}
	err = json.Unmarshal(tk, Tooken)
	if err != nil {
		return errors.New("WxPush:Token json.Unmarshal error")
	}
	if Tooken.Errcode != 0 {
		return fmt.Errorf("WxPush:Token Get error, %s", Tooken.Errmsg)
	}
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", Tooken.AccessToken)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.New("WxPush:json Marshal error")
	}
	post, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return errors.New("WxPush:Message Send error")
	}
	defer post.Body.Close()
	req, err := io.ReadAll(post.Body)
	if err != nil {
		return errors.New("WxPush:Push io.ReadAll error")
	}
	resData := &res{}
	err = json.Unmarshal(req, resData)
	if err != nil {
		return errors.New("WxPush:Push json.Unmarshal error")
	}
	if resData.Errcode != 0 {
		log.Error("WxPush: ", resData.Errcode, resData.Errmsg)
		return fmt.Errorf("WxPush: %v, %s", resData.Errcode, resData.Errmsg)
	}
	return nil
}
