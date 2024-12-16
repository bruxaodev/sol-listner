package connection

import (
	"encoding/json"
	"time"

	"github.com/bruxaodev/sol-listner/pkg/config"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type RpcInterface interface {
	Connection(url string, subscribe Subscribe, c chan []byte)
}

type Rpc struct {
	Socket *websocket.Conn
}

func NewRpc() RpcInterface {
	return &Rpc{
		Socket: nil,
	}
}

func (r *Rpc) Connection(url string, subscribe Subscribe, c chan []byte) {
	if r.Socket != nil {
		r.Socket.Close()
		r.Socket = nil
	}

	socket, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		config.Logger.Error("dial", zap.Error(err))
	}
	defer socket.Close()
	r.Socket = socket
	config.Logger.Info("connected", zap.String("url", url))

	subscriber, _ := json.Marshal(subscribe)
	socket.WriteMessage(websocket.TextMessage, subscriber)

	//mantain connection
	go func() {
		for {
			if socket != nil {
				socket.WriteMessage(websocket.PingMessage, []byte{})
			}
			time.Sleep(10 * time.Second)
		}
	}()

	for {
		_, message, err := socket.ReadMessage()
		if err != nil {
			config.Logger.Error("read", zap.Error(err))
			return
		}
		c <- message
	}
}
