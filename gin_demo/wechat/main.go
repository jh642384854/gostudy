package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/base"
	"github.com/chanxuehong/wechat/mp/material"
	"log"
	"net/http"

	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
)

const (
	wxAppId     = "wxf72825022bab4655"
	wxAppSecret = "622a90e691a1e2cfd98b7ccbdebce6a5"

	wxOriId         = "gh_fbceb886a0e4"
	wxToken         = "jhgocms"
	wxEncodedAESKey = "M1ToTaW76pGTqwTcVrjwnBhIuzILxqd7LYqSAf5iQAF"
)

var (
	// 下面变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler core.Handler
	msgServer  *core.Server
	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(wxAppId,wxAppSecret,nil)
	wechatClient *core.Client = core.NewClient(accessTokenServer,nil)
)

func init() {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	msgHandler = mux
	msgServer = core.NewServer(wxOriId, wxAppId, wxToken, wxEncodedAESKey, msgHandler, nil)
}

func textMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)

	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func init() {
	http.HandleFunc("/wx", wxCallbackHandler)
	http.HandleFunc("/getcallip", func(w http.ResponseWriter, i *http.Request) {
		ips,err := base.GetCallbackIP(wechatClient)
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Println(ips)
	})
	http.HandleFunc("/getmaterial", func(w http.ResponseWriter, i *http.Request) {
		rslt,err := material.BatchGet(wechatClient,"image",0,10)
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Println("素材总数：",rslt.TotalCount,rslt.ItemCount)
		fmt.Println("rslt:",rslt)
		for _, item := range rslt.Items {
			fmt.Println(item.Name,item.URL,item.MediaId,item.UpdateTime)
		}
		//w.Write();
	})
}

// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func wxCallbackHandler(w http.ResponseWriter, r *http.Request) {
	msgServer.ServeHTTP(w, r, nil)
}

func main() {
	fmt.Println("wechat server")
	log.Println(http.ListenAndServe(":8080", nil))
}