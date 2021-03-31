package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

const IndexFile = "./public/index.html"

type die struct {
	sides int
	name  string
}

var dice = map[string]*die{
	"d4": &die{4, "d4"},
}

func main() {
	r := gin.Default()
	m := melody.New()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Println(string(msg))
		m.Broadcast(msg)
	})

	r.Run("127.0.0.1:8080")
}
