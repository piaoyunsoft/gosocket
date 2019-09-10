package gosocket

import (
	"fmt"
	"net"
	"sync"
)

// TCPClient TCP客户端描述
type TCPClient struct {
	tcpClientState
	c           *Connection
	userHandler tcpEventHandler
	wg          sync.WaitGroup
}

type tcpClientState struct {
	remoteAddr string
	remotePort uint16
	connected  bool
}

// CreateTCPClient 创建一个TCPClient实例
func CreateTCPClient(handler tcpEventHandler) *TCPClient {
	return &TCPClient{userHandler: handler}
}

// Connect 连接到服务器
func (client *TCPClient) Connect(addr string, port uint16) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}
	client.c = newConnection(conn)

	// Fixed By PiaoYun -- 连接回调 2017-07-29
	if client.userHandler != nil {
		client.userHandler.OnConnect(client.c)
	}

	client.tcpClientState = tcpClientState{
		remoteAddr: addr,
		remotePort: port,
		connected:  true,
	}

	client.wg.Add(2)
	go client.c.recvThread(&client.wg, client.userHandler)
	go client.c.sendThread(&client.wg)
	return nil
}

// Send 发送数据
func (client *TCPClient) Send(data []byte) {
	client.c.Send(data)
}

// Close 关闭连接
func (client *TCPClient) Close() {
	client.c.Close()
	client.wg.Wait()
	if client.userHandler != nil {
		client.userHandler.OnDisconnect(client.c)
	}
}

// RemoteAddr 返回服务器地址
func (client *TCPClient) RemoteAddr() string {
	return fmt.Sprintf("%s:%d", client.tcpClientState.remoteAddr, client.tcpClientState.remotePort)
}

// LocalAddr 返回本机的连接地址
func (client *TCPClient) LocalAddr() string {
	return client.c.conn.LocalAddr().String()
}
