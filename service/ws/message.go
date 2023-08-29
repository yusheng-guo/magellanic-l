package ws

import (
	"time"
)

type MessageType uint8

const (
	MessageTypeRegister  MessageType = iota + 1 // 注册
	MessageTypeLogout                           // 注销
	MessageTypeHeartbeat                        // 心跳
	MessageTypeOneToOne                         // 1 => 1                         //
	MessageTypeChannel                          // 频道消息
	MessageTypeBroadcast                        // 广播消息
	MessageTypeEcho                             // 广播消息
)

// Message websocket 消息
type Message struct {
	Type      MessageType `json:"type"`      // 消息类型
	Content   []byte      `json:"content"`   // 消息内容
	From      string      `json:"from"`      // 来源
	To        string      `json:"to"`        // 去向
	Timestamp int64       `json:"timestamp"` // 时间戳
}

// NewMessage 实例化一条消息
func NewMessage(t MessageType, data []byte, from, to string) Message {
	return Message{
		Type:      t,
		Content:   data,
		From:      from,
		To:        to,
		Timestamp: time.Now().Unix(),
	}
}
