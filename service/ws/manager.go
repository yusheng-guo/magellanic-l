package ws

import (
	"github.com/yushengguo557/magellanic-l/global"
	"go.uber.org/zap"
)

// WebSocketManager websocket管理器
type WebSocketManager struct {
	Clients  map[string]*Client
	Channels map[string]*Channel
	Messages chan Message
}

// NewWebSocketManager 实例化websocket管理器
// cap: 消息通道容量
func NewWebSocketManager(cap int) *WebSocketManager {
	return &WebSocketManager{
		Clients:  make(map[string]*Client),
		Channels: make(map[string]*Channel),
		Messages: make(chan Message, cap),
	}
}

// Register 使用 WebSocketManager 对 client 进行管理 & 接收客户端发送过来的所有消息
func (m *WebSocketManager) Register(client *Client) {
	go func() {
		for {
			msg, err := client.Read()
			if err != nil {
				global.App.Log.Error("read data from client", zap.Any("err", err))
				continue
			}
			m.Messages <- msg
		}
	}()
	m.Clients[client.UID] = client
}

// Logout 取消 WebSocketManager 对 client 的管理 & 从所有频道中移除该客户端
func (m *WebSocketManager) Logout(uid string) {
	delete(m.Clients, uid)

	for _, ch := range m.Channels {
		delete(ch.Members, uid)
	}
}

// Broadcast 广播消息
func (m *WebSocketManager) Broadcast(msg Message) (err error) {
	for _, c := range m.Clients {
		err = c.Write(msg)
	}
	return err
}

// ReceiveMessage 接收消息
func (m *WebSocketManager) ReceiveMessage() {
	//for  {
	//
	//}
}

// HandlerMessage 处理消息
func (m *WebSocketManager) HandlerMessage() {
	var err error
	for msg := range m.Messages {
		switch msg.Type {
		case MessageTypeRegister:
			if m.IsManaged(msg.From) {
				//m.Register(m.Clients[msg.From])
			} else {
				m.PushMessage(msg)
			}
		case MessageTypeLogout:
			if m.IsManaged(msg.From) {
				m.Logout(msg.From)
			} else {
				m.PushMessage(msg)
			}
		case MessageTypeHeartbeat:
			echo := NewMessage(MessageTypeHeartbeat, []byte("success"), "", msg.From)
			err = m.Clients[msg.From].Write(echo)
		case MessageTypeOneToOne:
			// TODO: 私聊
		case MessageTypeChannel:
			err = m.Channels[msg.To].Write(msg)
		case MessageTypeBroadcast:
			err = m.Broadcast(msg)
		}
		if err != nil {
			global.App.Log.Error(err.Error())
		}
	}
}

func (m *WebSocketManager) PushMessage(msg Message) {
	m.Messages <- msg
}

// IsManaged 判断用户是否被 websocket 管理器 所管理
func (m *WebSocketManager) IsManaged(uid string) bool {
	_, ok := m.Clients[uid]
	return ok
}
