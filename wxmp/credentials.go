package wxmp

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"time"
	mu "../myutils"
)

type TokenReqParam struct {
	grant_type string
	appid string
	secret string
}

type Credential struct {
    appid string
	secret string
	token string
}

var _token chan string
var _credential = Credential{}

func TokenServer(appid, secret string) {
    tokParam := TokenReqParam{"client_credential", appid, secret}
    tokenServer(&tokParam)
}
func tokenServer(tokParam *TokenReqParam) {
    fmt.Println("11111111111")
	param := mu.MakeHttpGetParamStr(tokParam)
	var msg map[string]interface{}
	_token = make(chan string)
	_credential = Credential{tokParam.appid, tokParam.secret, ""}
	go func(){_token <- "ssss"}()

	fmt.Println("asfasdfweafsaf")
	go func(){
		// request for token
		go func(){
    		for{
    			<-_token
    			resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?" + param)
    			defer resp.Body.Close()
                body, err := ioutil.ReadAll(resp.Body)
                fmt.Println(string(body))
                err = json.Unmarshal(body, &msg)
                if err != nil {
                        go func(){_token <- "ssss"}()
                        time.Sleep(time.Second * 2)
                } else{
                    fmt.Printf("\ntoken%#v, %#v\n", msg["access_token"], msg["expires_in"])
                    if(msg["access_token"] == nil) {
                        go func(){_token <- "ssss"}()
                    } else {
                        _credential.token = fmt.Sprintf("%v", msg["access_token"])
                    }
                }
    		}
    		close(_token)
		}()
    	// refresh every 1.5hr
    	go func(){
    		for{
    			select{
    			case <- time.After(1.5 * 60 * time.Minute):
    				_token <- _credential.token
    			    fmt.Println("timeout: token")
    			}
    		}
    	}()
	}()
}

func GetToken() string{
    return _credential.token
}

func GetCredential() Credential{
    return _credential
}