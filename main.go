package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	app := gin.Default()


	app.GET("/", func(c *gin.Context) {
	  c.String(http.StatusOK, "Olá, visitante, eu estou rodando em Golang!")
	})

	app.GET("/decola", func(c *gin.Context) {
		c.String(http.StatusOK, "Hmm... Parece que você faz parte do decola, né?")
	  })

	  app.GET("/:nome", func(c *gin.Context) {
		name := c.Param("nome")
		c.String(http.StatusOK, "Olá, %s! Sabia que estou rodando em Golang?", name)
	  })


	app.Run(":80") 
  }
