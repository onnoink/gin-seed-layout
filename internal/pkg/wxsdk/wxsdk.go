package wxsdk

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	appId     string
	appSecret string
	sync.Mutex
}

func NewClient(appid string, appSecret string) *Client {
	return &Client{
		appId:     appid,
		appSecret: appSecret,
	}
}

type Code2SessionReply struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type GetAccessTokenReply struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int64  `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func (c *Client) Code2Session(code string) (*Code2SessionReply, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", c.appId, c.appSecret, code)
	request := gorequest.New().Timeout(2 * time.Second)
	resp, body, errs := request.Get(url).End()
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("请求微信服务器错误,返回状态码错误")
	}

	reply := new(Code2SessionReply)
	if err := json.Unmarshal([]byte(body), reply); err != nil {
		return nil, errors.New("解析微信服务器返回值失败")
	}

	if reply.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("请求微信服务器错误,返回状态码错误%s", reply.ErrMsg))
	}

	return reply, nil
}
