package tool

import (
	"encoding/json"
	"errors"
	"fmt"
	email_ "github.com/jordan-wright/email"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/smtp"
	"net/textproto"
	"qiublog/utils"
	"strings"
)

var push = utils.Config.Server.Push

func Send(title, content string) error {
	switch strings.ToUpper(push.Enable) {
	case "WXPUSH":
		return wxPush(title, content)
	case "EMAIL":
		return email(title, content)
	}
	return nil
}

func wxPush(title, content string) error {
	type body struct {
		Agentid string `json:"agentid"`
		Touser  string `json:"touser"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
		Msgtype string `json:"msgtype"`
	}
	Agentid, CorpID, Secret := push.WxPush.Agentid, push.WxPush.CorpId, push.WxPush.Secret
	data := body{
		Agentid: Agentid,
		Touser:  "@all",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: title + content,
		},
		Msgtype: "text",
	}
	tookenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", CorpID, Secret)
	get, err := http.Get(tookenUrl)
	if err != nil {
		return errors.New("WxPush:Token Get error")
	}
	defer get.Body.Close()
	tk, err := io.ReadAll(get.Body)
	if err != nil {
		return errors.New("WxPush:Token io.ReadAll error")
	}
	Tooken := &struct {
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}{}
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
	resData := &struct {
		Errcode        int    `json:"errcode"`
		Errmsg         string `json:"errmsg"`
		Invaliduser    string `json:"invaliduser"`
		Invalidparty   string `json:"invalidparty"`
		Invalidtag     string `json:"invalidtag"`
		Unlicenseduser string `json:"unlicenseduser"`
		Msgid          string `json:"msgid"`
		ResponseCode   string `json:"response_code"`
	}{}
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

func email(title, content string) error {
	tos, password, from, host, port := push.Email.To, push.Email.Password, push.Email.From, push.Email.Host, push.Email.Port
	var to []string
	if tos == "" {
		to = []string{from}
	} else {
		to = strings.Split(tos, ",")
	}
	e := &email_.Email{
		To:      to,
		From:    fmt.Sprintf("QiuBlogBot <%s>", from),
		Subject: title,
		Text:    []byte(content),
		Headers: textproto.MIMEHeader{},
	}
	return e.Send(host+":"+port, smtp.PlainAuth("", from, password, host))
}
