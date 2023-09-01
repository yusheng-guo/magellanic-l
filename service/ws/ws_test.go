package ws

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	origin := "http://127.0.0.1:9999/"
	url := "ws://127.0.0.1:9999/api/v1/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	go receiveMessage(ws)

	sendMessage(ws, MessageTypeRegister, "register", "")
	for i := 0; i < 5; i++ {
		sendMessage(ws, MessageTypeHeartbeat, "heartbeat", "")
		time.Sleep(3 * time.Second)
	}
	sendMessage(ws, MessageTypeLogout, "logout", "")
}

func sendMessage(ws *websocket.Conn, t MessageType, text string, to string) {
	var send Message
	var err error
	send = NewMessage(t, []byte(text), "", to)
	fmt.Printf("send: %#v\n", string(send.Content))
	err = websocket.JSON.Send(ws, send)
	if err != nil {
		panic(err)
	}
}

func receiveMessage(ws *websocket.Conn) {
	var recv Message
	var err error
	for {
		err = websocket.JSON.Receive(ws, &recv)
		if err != nil {
			panic(err)
		}
		fmt.Printf("receive: %#v\n", string(recv.Content))
	}
}
