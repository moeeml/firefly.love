package lib

import (
	"encoding/json"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/websocket"
)

func MapToMessage(msg context.Map) websocket.Message{
	json, _ := json.Marshal(msg)
	message := websocket.Message{
		Body:     json,
		IsNative: true,
	}

	return message
}