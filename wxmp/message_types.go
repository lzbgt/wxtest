package wxmp

import (
	"encoding/xml"
	"time"
	mu "myapp/myutils"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type WildRequest struct {
	// base
    XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
    MsgId 	     uint64
    Encrypt      string

    // text
    Content 	 string


    // base media
    MediaId      string

    // pic
    PicUrl       string

    // voice
    Format       string

    // vr
    Recognition  string

    // video
    ThumbMediaId string

    // position
    Location_X   string
    Location_Y   string
    Scale        int
    Label        string

    // link
    Title        string
    Description  string
    Url          string

    // base event
    Event        string

    // QR Code
    EventKey     string
    Ticket       string

    // position report
    Latitude     string
    Longitude    string
    Precision    float32

    // menu click

    // menu travel
}

type BaseRequest struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	MsgId 	     uint64
	Event        string
	EventKey     string
}

type BaseResponse struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
}

type TxtRequest struct {
	BaseRequest
	Content string
	MsgId   string
}

type LocRequest struct {
	BaseRequest
	Location_X float64
	Location_Y float64
	Scale      int32
	Label      string
}

type TxtResponse struct {
	XMLName xml.Name `xml:"xml"`
	BaseResponse
	Content string
}

type PicResponse struct {
	XMLName xml.Name `xml:"xml"`
	BaseResponse
	ArticleCount int
	Articles     *Articles
}

type Item struct {
	Title       string
	Description string
	PicUrl      string
	Url         string
}

type Articles struct {
	// Articles xml.Name `xml:"Articles"`
	Items []*Item `xml:"item"`
}

type  UnionIDReq struct {
    access_token string
    openid       string
    lang         string  //zh_CN, zh_TW, en
}

func (ui *UnionIDReq) Get(base *string) interface{} {
    if base == nil {
        base = new(string)
        *base = "https://api.weixin.qq.com/cgi-bin/user/info?"
    }
    param := mu.MakeHttpGetParamStr(ui)
	resp, _ := http.Get(*base + param)
	defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    var msg map[string]interface{}
    _ = json.Unmarshal(body, &msg)
    return msg
}
