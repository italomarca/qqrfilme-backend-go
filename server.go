package main

import (
	"github.com/gin-gonic/gin"
	movies "github.com/italomarca/qqrfilme-backend-go/www"
)

func main() {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/movie", movies.GetRandomMovie)
	}
	router.Run()
}
