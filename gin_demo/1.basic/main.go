package main

import (
	"github.com/chanxuehong/wechat/mp/base"
	"fmt"
	"github.com/chanxuehong/wechat/mp/material"
	"log"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
	入门基本示例
 */
var (
	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer("wxf72825022bab4655","622a90e691a1e2cfd98b7ccbdebce6a5",nil)
	wechatClient *core.Client = core.NewClient(accessTokenServer,nil)
)
func main() {
	r := gin.Default()
	data := map[string]interface{}{
		"title":"gin 学习",
		"days":15,
	}
	//1.入门示例
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"pong",
		})
	})

	//2.AsciiJSON：使用AsciiJSON生成只有ascii的JSON，并使用转义的非ascii计时器。这个会输出{"days":15,"title":"gin \u5b66\u4e60"}
	r.GET("/asciiiJson", func(context *gin.Context) {
		context.AsciiJSON(http.StatusOK,data)
	})
	//3.JSON：这个输出会被转换为：{"html":"\u003cb\u003eHello Gin\u003c/b\u003e"}
	r.GET("/normaljson", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"html":"<b>Hello Gin</b>",
		})
	})
	//4.PureJSON：这个会原样输出：{"html":"<b>Hello Gin</b>"}
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK,gin.H{
			"html":"<b>Hello Gin</b>",
		})
	})

	//5.SecureJSON：将输出while(1);["go","java","php"]
	r.GET("/secureJson", func(c *gin.Context) {
		names := []string{"go","java","php"}
		c.SecureJSON(http.StatusOK,names)
	})

	//6.IndentedJSON：这种会带格式的输出，以下会输出以下内容：
	/**
		{
			"days": 15,
			"title": "gin 学习"
		}
	 */
	r.GET("/indentedJson", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK,data)
	})
	//要转换为xml格式，需要特定的对象
	r.GET("/xmldata", func(c *gin.Context) {
		c.XML(http.StatusOK,data)
	})

	r.GET("/yamldata", func(c *gin.Context) {
		c.YAML(http.StatusOK,data)
	})
	//这个需要特殊处理。
	r.GET("/protobuf", func(c *gin.Context) {
		c.ProtoBuf(http.StatusOK,data)
	})

	r.GET("/stringdata", func(c *gin.Context) {
		c.String(http.StatusOK,"%v:%d","age",15)
	})

	/**
		获取微信服务器IP地址
	 */
	r.GET("/getcallip", func(c *gin.Context) {
		ips,err := base.GetCallbackIP(wechatClient)
		if err != nil{
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK,ips)
	})
	/**
		获取永久素材列表。图片（image）、视频（video）、语音 （voice）
	 */
	r.GET("/getmaterial", func(c *gin.Context) {
		//fmt.Println(wechatClient.Token())
		//materialType参数值可以是：图片（image）、视频（video）、语音 （voice）、图文（news）
		rslt,err := material.BatchGet(wechatClient,"video",0,10)
		if err != nil{
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK,rslt)
	})
	/**
		获取图文素材列表。
	 */
	r.GET("/getnewsmaterial", func(c *gin.Context) {
		//materialType参数值可以是：图片（image）、视频（video）、语音 （voice）、图文（news）
		rslt,err := material.BatchGetNews(wechatClient,0,10)
		if err != nil{
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK,rslt)
	})

	/**
		获取某个素材信息
	 */
	r.GET("/getonematerial", func(c *gin.Context) {
		video,err := material.GetVideo(wechatClient,"8tm_L653x9N30xEwyRpqDiMgcKsPb3i0xPsjsKY5ihY")
		if err != nil{
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK,video)
	})


	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	//微信服务器
	wxServer := core.NewServer("gh_fbceb886a0e4","wxf72825022bab4655","jhgocms","M1ToTaW76pGTqwTcVrjwnBhIuzILxqd7LYqSAf5iQAF",mux,nil)

	r.Any("/wx", func(c *gin.Context) {
		wxServer.ServeHTTP(c.Writer,c.Request,nil)
	})
	r.Run()
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