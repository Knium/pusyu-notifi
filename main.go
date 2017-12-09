package main

import (
	"github.com/labstack/echo"
	"net/http"
	"os/exec"
	"gopkg.in/olahol/melody.v1"

)

func main()  {
	server := echo.New()
	m := melody.New()
	server.GET("/", func(c echo.Context) error {
		http.ServeFile(c.Response(), c.Request(), "index.html")
		return nil
	})
	server.GET("/echo", func(context echo.Context) error {
		e := exec.Command("sh", "-c", "echo Hello, World! > /dev/ttys000")
		e.Start()
		return context.String(http.StatusOK, "done")
	})
	server.GET("/ws", func(c echo.Context) error {
		m.HandleRequest(c.Response(), c.Request())
		return nil
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	server.Start(":8002")
}