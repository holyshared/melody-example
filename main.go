package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"
  "github.com/holyshared/melody-example/app"
)

func main() {
  server := gin.Default()
  websocket := app.NewApp()

  server.Use(static.Serve("/", static.LocalFile("./public", true)))
  server.GET("/ws", websocket.Handler)
  server.Run(":3000")
}
