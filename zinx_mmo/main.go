package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
	"zinx_mmo/core"
)

func OnConnectionAdd(connection ziface.IConnection)  {
	player := core.NewPlayer(connection)
	fmt.Println("=====> Player pid = ", player.Pid, " is arrived <=====")
}

func main() {

	server := znet.NewServer("mmo test")

	server.SetOnConnectStart(OnConnectionAdd)

	server.Run()
}
