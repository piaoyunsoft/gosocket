# gosocket
一个轻量级的Go语言 socket通信库, 与协议无关，简洁好用的API接口

基于<https://github.com/xikug/gsocket/>做了些修改

## 0x00 安装

``go get -u github.com/piaoyunsoft/gosocket``


## 0x01 服务端

```
package main

import (
	"bufio"
	"log"
	"os"
	"github.com/piaoyunsoft/gosocket"
)

// 实现tcpEventHandler接口即可
type demoServer struct{}

func (server demoServer) OnConnect(c *gsocket.Connection) {
	log.Printf("CONNECTED: %s\n", c.RemoteAddr())
}

func (server demoServer) OnDisconnect(c *gsocket.Connection) {
	log.Printf("DISCONNECTED: %s\n", c.RemoteAddr())
}

func (server demoServer) OnRecv(c *gsocket.Connection, data []byte) {
	log.Printf("DATA RECVED: %s %d - %v\n", c.RemoteAddr(), len(data), data)
	c.Send(data)
}

func (server demoServer) OnError(c *gsocket.Connection, err error) {
	log.Printf("ERROR: %s - %s\n", c.RemoteAddr(), err.Error())
}

func main() {
	demoServer := &demoServer{}

	//CreateTCPServer 的handler可以传nil
	server := gsocket.CreateTCPServer("0.0.0.0", 1234, demoServer)

	err := server.Start()
	if err != nil {
		log.Printf("Start Server Error: %s\n", err.Error())
		return
	}

	log.Printf("Listening %s...\n", server.Addr())

	select()
}

```

## 0x2 客户端

```
package main

import (
	"bufio"
	"log"
	"os"
	"github.com/piaoyunsoft/gosocket"
)

// 实现tcpEventHandler接口即可
type demoClient struct{}

func (client *demoClient) OnConnect(c *gsocket.Connection) {
	log.Printf("CONNECTED: %s\n", c.RemoteAddr())
}

func (client *demoClient) OnDisconnect(c *gsocket.Connection) {
	log.Printf("DISCONNECTED: %s\n", c.RemoteAddr())
}

func (client *demoClient) OnRecv(c *gsocket.Connection, data []byte) {
	log.Printf("DATA RECVED: %s %d - %v\n", c.RemoteAddr(), len(data), data)
}

func (client *demoClient) OnError(c *gsocket.Connection, err error) {
	log.Printf("ERROR: %s - %s\n", c.RemoteAddr(), err.Error())
}

func main() {
	demoClient := &demoClient{}

	client := gsocket.CreateTCPClient(demoClient)

	err := client.Connect("127.0.0.1", 1234)
	if err != nil {
		log.Printf("Coneect Server Error: %s\n", err.Error())
		return
	}

	log.Printf("Connect Server %s Success\n", client.RemoteAddr())

	client.Send([]byte("Hello PiaoPiao(https://www.chinapyg.com)"))

	client.Close()
}
```