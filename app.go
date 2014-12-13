package main

import (
	"fmt"
	"./wxmp"
	"net/http"
)

func main() {
    wxmp.TokenServer("wxd01846ec94dbe0d6", "e33b2fbf1b26762de92d2f87a5059aff")
    http.HandleFunc("/weixin/access", MsgHandler)
    http.ListenAndServe(":80", nil)
}


func MsgHandler(w http.ResponseWriter, r *http.Request) {
    var mp = wxmp.New(&w, r)
    fmt.Printf("%#v\n", mp)

    if mp.Err != 0 {
        return
    }

	if "event" == mp.MsgType && mp.Event == "subscribe" {
		fmt.Println("do subscribe")
	} else if "event" == mp.MsgType && mp.Event == "unsubscribe" {
	    fmt.Println("do unsubscribe")
	} else if "location" == mp.MsgType {
	    fmt.Println("do location")
	} else if "text" == mp.MsgType {
	    fmt.Println("text msg")
	    mp.ReplyTextMsg("Hello")
	    return
	} else if "image" == mp.MsgType {
	    fmt.Println("image msg")
	} else if "voice" == mp.MsgType {
	    fmt.Println("voice msg")
	} else if "video" == mp.MsgType {
	    fmt.Println("video msg")
	} else if "location" == mp.MsgType {
	    fmt.Println("location msg")
	} else if "link" == mp.MsgType {
	    fmt.Println("link msg")
	} else if "mp.Event" == mp.MsgType && mp.Event == "SCAN" {
	    fmt.Println("SCAN mp.Event")
	} else if "mp.Event" == mp.MsgType && mp.Event == "LOCATION" {
	    fmt.Println("LOCATION mp.Event")
	} else if "mp.Event" == mp.MsgType && mp.Event == "CLICK" {
	    fmt.Println("CLICK mp.Event")
	} else if "mp.Event" == mp.MsgType && mp.Event == "VIEW" {
	    fmt.Println("VIEW mp.Event")
	} else if "mp.Event" == mp.MsgType && mp.Event == "unsubscribe" {
	    fmt.Println("text mp.Event")
	} else {
	    fmt.Println("unknown msg")
	}
	mp.ReplyEmptyMsg()
}
