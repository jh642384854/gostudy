package main

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)
/**
	这个是服务端接收消息的结构体，在客户端向服务器发送数据的时候，客户端也需要按照这样的格式进行封装，然后在发送数据到服务端。
	客户端封装的数据需要用下面字段标签的属性名进行封装
 */
type Message struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Msg string `json:"msg"`
}

type Client struct {
	id string
	conn *websocket.Conn
}

var (
	//定义客户端信息
	clients = make(map[*websocket.Conn]*Client)
	//定义消息承载体
	braodcast = make(chan Message)
	//定义upgrader
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//ws服务
func wsServer(w http.ResponseWriter,r *http.Request)  {
	ws,err := upgrader.Upgrade(w,r,nil)
	if err != nil{
		log.Fatal(err)
	}
	defer ws.Close()
	uid,_:= uuid.NewV4()
	userid := uid.String()
	client := &Client{
		id:userid,
		conn:ws,
	}
	log.Println("当前客户端id：",userid)
	//注册客户端
	clients[ws] = client

	for  {
		var msg Message
		err := ws.ReadJSON(&msg)  //将获取的数据写到上面定义的变量中
		log.Println("read data from client")
		if err != nil{
			log.Printf("error :%v",err)
			delete(clients,ws)
			break
		}
		braodcast <- msg
	}
}
//处理消息
func handleMessages()  {
	for  {
		msg := <- braodcast
		log.Println("send data to client")
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil{
				log.Printf("error : %v",err)
				client.Close()
				delete(clients,client)
			}
		}
	}
}

func main() {
	server := http.Server{
		Addr:":8090",
	}
	fs := http.FileServer(http.Dir("E:/GoProjects/src/dev/websocketdemo/public"))
	http.Handle("/",fs)
	http.HandleFunc("/ws",wsServer)

	go handleMessages()

	log.Println("http server start on : 8090")
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}
