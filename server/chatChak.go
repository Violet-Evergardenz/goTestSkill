package server

import (
	// "log"
	"fmt"
	so "myapp/server/service/utilities/socket"

	socketio "github.com/googollee/go-socket.io"

	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *FuncHandler) JoinChatRoom(c echo.Context) error {
	so.Server.OnConnect("/boss", func(c socketio.Conn) error {
		c.Join("/outofbrain")
		c.Emit("chat", "msg")
		fmt.Println("Boss connected")
		return nil
	})
	return c.JSON(http.StatusOK, "")
}
