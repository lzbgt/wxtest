package wxmp

import (
	"bytes"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

type WxMP struct {
	*WildRequest
	resp *http.ResponseWriter
	Err  int
}

const (
	//
	CDATA_OPEN        = "<![CDATA["
	CDATA_ALIAS_OPEN  = "!+!"
	CDATA_CLOSE       = "]]"
	CDATA_ALIAS_CLOSE = "!x!"

	//
	E_OK         = 0
	E_ACCESS     = 1
	E_READ_BODY  = 2
	E_PARSE_BODY = 3
	E_UNKNOWN    = 5
)

func New(resp *http.ResponseWriter, req *http.Request) *WxMP {
	mp := &WxMP{nil, resp, E_UNKNOWN}
	data, err := ioutil.ReadAll(req.Body)
	if nil != err {
		fmt.Println("read body error:", err)
		mp.Err = E_READ_BODY
		return nil
	}
	fmt.Println(string(data))

	if req.Method == "GET" {
		AccessVerify(resp, req)
		mp.Err = E_ACCESS
		return nil
	} else {
		request := new(WildRequest)
		er := xml.Unmarshal(data, request)
		if nil != er {
			fmt.Println("parse body err:", er)
			mp.Err = E_PARSE_BODY
			return nil
		}
		fmt.Printf("\n%#v\n", request)
		mp.WildRequest = request
		mp.Err = E_OK
	}
	return mp
}

func (this *WxMP) ReplyEmptyMsg() {
	if _, err := (*this.resp).Write([]byte("")); err != nil {
		fmt.Println("\n\nError reply\n")
	}
}

func (this *WxMP) ReplyTextMsgEscape(content string) {
	txtMsg := TxtResponse{}
	txtMsg.ToUserName = EscapWXText(this.FromUserName)
	txtMsg.FromUserName = EscapWXText(this.ToUserName)
	txtMsg.CreateTime = time.Duration(time.Now().Unix())
	txtMsg.MsgType = EscapWXText("text")
	txtMsg.Content = EscapWXText(content)
	brespons, _ := xml.MarshalIndent(txtMsg, "", "    ")
	brespons = bytes.Replace(brespons, []byte(CDATA_ALIAS_OPEN), []byte(CDATA_OPEN), -1)
	brespons = bytes.Replace(brespons, []byte(CDATA_ALIAS_CLOSE), []byte(CDATA_CLOSE), -1)
	fmt.Println(string(brespons))
	(*this.resp).Header().Set("Content-Type", "text/xml")
	if _, err := (*this.resp).Write(brespons); err != nil {
		fmt.Println("\n\nError reply\n")
	}
}

func (this *WxMP) ReplyTextMsg(content string) {
	txtMsg := TxtResponse{}
	txtMsg.ToUserName = this.FromUserName
	txtMsg.FromUserName = this.ToUserName
	txtMsg.CreateTime = time.Duration(time.Now().Unix())
	txtMsg.MsgType = "text"
	txtMsg.Content = content
	brespons, _ := xml.Marshal(txtMsg)
	fmt.Println(string(brespons))
	(*this.resp).Header().Set("Content-Type", "text/xml")
	if _, err := (*this.resp).Write(brespons); err != nil {
		fmt.Println("\n\nError reply\n")
	}
}

func EscapWXText(s interface{}) string {
	if str, ok := s.(string); ok {
		return CDATA_ALIAS_OPEN + str + CDATA_ALIAS_CLOSE
	}

	if str, ok := s.([]byte); ok {
		var b bytes.Buffer
		b.Write([]byte(CDATA_ALIAS_OPEN + string(str) + CDATA_ALIAS_CLOSE))
		return b.String()
	}
	str := "INVALID"
	fmt.Println(str)
	return str
}

func AccessVerify(resp *http.ResponseWriter, req *http.Request) bool {
	signature := req.FormValue("signature")
	timestamp := req.FormValue("timestamp")
	nonce := req.FormValue("nonce")
	echostr := req.FormValue("echostr")

	fmt.Printf("\n%s, %s, %s, %s\n", signature, timestamp, nonce, echostr)

	sign := func() string {
		strs := sort.StringSlice{"lzbgt1", timestamp, nonce}
		sort.Strings(strs)
		str := ""
		for _, s := range strs {
			str += s
		}

		h := sha1.New()
		h.Write([]byte(str))
		return fmt.Sprintf("%x", h.Sum(nil))
	}()

	fmt.Printf("\ncalc %s, rcv %s\n", sign, signature)

	if sign == signature {
		(*resp).Write([]byte(echostr))
	} else {
		(*resp).Write([]byte(""))
		return false
	}
	return true
}
