package gosocket

import (
	"bufio"
	"log"
	"os"
	"testing"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type demoServer struct{}

// demoServer 服务端

func (server demoServer) OnConnect(c *Connection) {
	log.Printf("CONNECTED: %s\n", c.RemoteAddr())
}

func (server demoServer) OnDisconnect(c *Connection) {
	log.Printf("DISCONNECTED: %s\n", c.RemoteAddr())
}

func (server demoServer) OnRecv(c *Connection, data []byte) {
	log.Printf("DATA RECVED: %s %d - %v\n", c.RemoteAddr(), len(data), string(data))
	c.Send(append([]byte(">"), data...))
}

func (server demoServer) OnError(c *Connection, err error) {
	log.Printf("ERROR: %s - %s\n", c.RemoteAddr(), err.Error())
}

func serverStart() {
	demoServer := &demoServer{}
	// CreateTCPServer 的handler可以传nil
	server := CreateTCPServer("0.0.0.0", 1234, demoServer)

	err := server.Start()
	if err != nil {
		log.Printf("Start Server Error: %s\n", err.Error())
		return
	}

	log.Printf("Listening %s...\n", server.Addr())

	select {}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// demoClient 客户端
type demoClient struct{}

func (client *demoClient) OnConnect(c *Connection) {
	log.Printf("CONNECTED: %s\n", c.RemoteAddr())
}

func (client *demoClient) OnDisconnect(c *Connection) {
	log.Printf("DISCONNECTED: %s\n", c.RemoteAddr())
}

func (client *demoClient) OnRecv(c *Connection, data []byte) {
	log.Printf("DATA RECVED: %s %d - %v\n", c.RemoteAddr(), len(data), string(data))
}

func (client *demoClient) OnError(c *Connection, err error) {
	log.Printf("ERROR: %s - %s\n", c.RemoteAddr(), err.Error())
}

func clientStart() {
	demoClient := &demoClient{}

	client := CreateTCPClient(demoClient)

	err := client.Connect("127.0.0.1", 1234)
	if err != nil {
		log.Printf("Coneect Server Error: %s\n", err.Error())
		return
	}

	log.Printf("Connect Server %s Success\n", client.RemoteAddr())

	for {
		client.Send([]byte("Hello PiaoPiao(https://WwW.ChinaPYG.CoM)!!!"))
		time.Sleep(2 * time.Second)
	}

	client.Close()
}

func pause() {
	println("按回车键退出...\n")
	r := bufio.NewReader(os.Stdin)
	r.ReadByte()
}

func TestGoSocket(t *testing.T) {
	go serverStart()
	go clientStart()
	select {}
}
