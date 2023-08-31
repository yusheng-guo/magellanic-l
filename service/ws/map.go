package ws

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

const WebSocketClientToServerMap = "websocket-client-to-server-map"

type ClientToServerMap struct {
	rdb *redis.Client
}

func (m *ClientToServerMap) Del(key string) {
	if val := m.rdb.HDel(context.Background(), WebSocketClientToServerMap, key).Val(); val != 1 {
		log.Printf("failed delete client [%s] in redis\n", key)
	}
}

func (m *ClientToServerMap) Set(key, val string) error {
	if err := m.rdb.HSet(context.Background(), WebSocketClientToServerMap, key, val).Err(); err != nil {
		return fmt.Errorf("set client [%s] = manager [%s], err: %s\n", key, val, err)
	}
	return nil
}

func (m *ClientToServerMap) Get(key string) (string, error) {
	val, err := m.rdb.HGet(context.Background(), WebSocketClientToServerMap, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errors.New(fmt.Sprintf("client [%s] not exist\n", key))
		} else {
			log.Fatalf("get manager of client [%s], err: %s\n", key, err)
		}
	}
	return val, nil
}
