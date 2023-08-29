package ws

import (
	"fmt"
	"golang.org/x/net/websocket"
)

// Client websocket客户端
type Client struct {
	UID  string          // 用户ID
	Conn *websocket.Conn // websocket连接
}

// Write 向客户端中写入信息
func (c *Client) Write(msg Message) error {
	err := websocket.Message.Send(c.Conn, msg)
	if err != nil {
		return fmt.Errorf("send message to [%#v], err: %w", c, err)
	}
	return nil
}

// Read 接收消息
func (c *Client) Read() (Message, error) {
	var msg Message
	err := websocket.Message.Receive(c.Conn, &msg)
	if err != nil {
		return Message{}, fmt.Errorf("receive from [%#v], err: %w", c, err)
	}
	return msg, nil
}

// JoinChannel 加入频道
func (c *Client) JoinChannel(ch *Channel) {
	ch.Members[c.UID] = c
}

// LeaveChannel 离开频道
func (c *Client) LeaveChannel(ch *Channel) {
	delete(ch.Members, c.UID)
}
