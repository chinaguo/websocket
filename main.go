package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	"sync"
)

func main() {
	app := iris.New()
	ws := websocket.New(websocket.Config{})

	conn := make(map[string]websocket.Connection)
	chatRoom := "room"
	mx := new(sync.Mutex)

	ws.OnConnection(func(c websocket.Connection) {

		c.On("chat", func(message string) {
			if message == "join" {
				c.Join(chatRoom)
		        mx.Lock()
		        conn[""] = c
		        mx.Unlock()
			}
		})
	})

	app.Run(iris.Addr(":8080"))
}
