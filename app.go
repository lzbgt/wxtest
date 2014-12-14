package main

import (
	"fmt"
	"myapp/wxmp"
	"net/http"
)

func main() {
	wxmp.TokenServer("wxb7d39933d3a607af", "c115d488b48d0b27c8a9b10605177cf7")
	http.HandleFunc("/weixin/access", MsgHandler)
	http.ListenAndServe(":80", nil)
}

func MsgHandler(w http.ResponseWriter, r *http.Request) {
	mp := wxmp.New(&w, r)
	if mp == nil {
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
