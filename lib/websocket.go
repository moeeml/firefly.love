package lib

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/websocket"
	"strconv"
)

func CreateWebsocket() context.Handler {
	type Message struct {
		Type     string `json:type`
		X        string `json:x`
		Y        string `json:y`
		Angle    string `json:angle`
		Momentum string `json:momentum`
		Message  string `json:message`
		Name     string `json:name`
	}

	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			var message Message
			var rspMsg context.Map
			json.Unmarshal(msg.Body, &message)

			switch message.Type {
			case "login":
				return nil
			case "update":
				angle, _ := strconv.ParseFloat(message.Angle, 32)
				momentum, _ := strconv.ParseFloat(message.Momentum, 32)
				x, _ := strconv.ParseFloat(message.X, 32)
				y, _ := strconv.ParseFloat(message.Y, 32)

				rspMsg = iris.Map{
					"id":         nsConn.Conn.ID(),
					"name":       message.Name,
					"life":       1,
					"authorized": false,
					"type":       "update",
					"x":          x,
					"y":          y,
					"angle":      angle,
					"momentum":   momentum,
				}
			case "message":
				rspMsg = iris.Map{
					"type":    "message",
					"id":      nsConn.Conn.ID(),
					"message": message.Message,
				}
			}

			bMsg := MapToMessage(rspMsg)
			nsConn.Conn.Server().Broadcast(nil, bMsg)
			return nil
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {
		msg := iris.Map{"type": "welcome", "id": c.ID()}
		r := MapToMessage(msg)
		c.Write(r)
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		msg := iris.Map{"type": "closed", "id": c.ID()}
		r := MapToMessage(msg)
		c.Server().Broadcast(nil, r)
	}

	ws.OnUpgradeError = func(err error) {
		fmt.Printf("Upgrade Error: %v", err)
	}

	return websocket.Handler(ws)
}
