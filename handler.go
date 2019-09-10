package gosocket

type tcpEventHandler interface {
	// 连接事件
	OnConnect(c *Connection)
	// 断开连接事件
	OnDisconnect(c *Connection)
	// 收到数据事件
	OnRecv(c *Connection, data []byte)
	// 有错误发生
	OnError(c *Connection, err error)
}