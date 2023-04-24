package socket

import (
	socketio "github.com/googollee/go-socket.io"
)

var Server *socketio.Server

func init() {
	Server := socketio.NewServer(nil)
	if Server == nil {
		panic(Server)
	}
}
