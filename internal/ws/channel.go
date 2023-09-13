package ws

import (
	"fmt"
)

// Channel websocket频道
type Channel struct {
	ID      string             // 群组ID
	Members map[string]*Client // 群组成员 UID => WebSocket客户端
}

// NewChannel 实例化频道
func NewChannel(id string) Channel {
	return Channel{
		ID:      id,
		Members: make(map[string]*Client),
	}
}

// Write 向频道中写入消息
func (c *Channel) Write(msg Message) (err error) {
	for _, member := range c.Members {
		err := member.Write(msg)
		if err != nil {
			err = fmt.Errorf("write to channel [%#v], err: %w", c, err)
			continue
		}
	}
	return err
}

//func (c *Channel) Deport(cl *Client) {
//	delete(c.Members, c.ID)
//}
