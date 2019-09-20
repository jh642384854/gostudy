package main

import (
	"fmt"
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"time"
)

const (
	writeWait  = 10 * time.Second
	pontWait   = 60 * time.Second
	pingPeriod = (pontWait * 9) / 10
	filePeriod = 10 * time.Second
)

//定义client管理器
type ClientManager struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

//定义每个连接到websocket的客户端信息
type Client struct {
	//客户端唯一ID
	id   string
	//客户端链接
	conn *websocket.Conn
	send chan []byte
}

//定义消息体信息
type Message struct {
	Send    string
	Content string
}

//初始化ClientManager
var (
	addr     = flag.String("addr", ":8080", "http service address")
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	manager = ClientManager{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
)

func (manager *ClientManager) start()  {
	for  {
		select {
		case client := <- manager.register:
			manager.clients[client] = true
			message,_ := json.Marshal(&Message{
				Content:"/A new socket has connected.",
			})
			manager.send(message,client)
		case client := <- manager.unregister:
			if _,ok := manager.clients[client]; ok{
				close(client.send)
				delete(manager.clients,client)
				message,_ := json.Marshal(&Message{
					Content: "/A socket has disconnected.",
				})
				manager.send(message,client)
			}
		case msg := <-manager.broadcast:
			for client := range manager.clients {
				select {
				case client.send <- msg:
				default:
					close(client.send)
					delete(manager.clients,client)
				}
			}
		}
	}
}
//管理客户端发送数据
func (manager *ClientManager) send(message []byte,ingore *Client)  {
	for client := range manager.clients {
		if client != ingore{
			client.send <- message
		}
	}
}

//读客户端传递的数据
func (c *Client) reader() {
	defer func() {
		manager.unregister <- c
		c.conn.Close()
	}()

	for  {
		_,mesage,error := c.conn.ReadMessage()
		if error != nil{
			manager.unregister <- c
			c.conn.Close()
			break
		}
		msg,_ := json.Marshal(&Message{
			Send:c.id,
			Content:string(mesage),
		})
		manager.broadcast <- msg
	}
}
//给客户端写数据
func (c *Client) writer()  {
	defer func() {
		c.conn.Close()
	}()
	for  {
		select {
		case msg,ok := <- c.send:
			if !ok{
				c.conn.WriteMessage(websocket.CloseMessage,[]byte{})
				return
			}
			c.conn.WriteMessage(websocket.TextMessage,msg)
		}
	}
}

//serverHome(项目主页)
func serverHome(w http.ResponseWriter,r *http.Request)  {
	log.Println(r.URL)
	if r.URL.Path != "/"{
		http.Error(w,"Not Found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	}
	http.ServeFile(w,r,"E:/GoProjects/src/dev/websocketdemo/index.html")
}

//websocket服务端
func serverWs(w http.ResponseWriter, r *http.Request) {
	//升级get请求为webSocket协议
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	uid := uuid.Must(uuid.NewV4())
	client := &Client{
		id:   uid.String(),
		conn: conn,
		send: make(chan []byte),
	}
	manager.register <- client

	go client.reader()
	go client.writer()
}

func main() {
	fmt.Println("Starting application... localhost:8080")
	//flag.Parse()
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/",serverHome)
	http.HandleFunc("/ws", serverWs)
	server.ListenAndServe()
}
