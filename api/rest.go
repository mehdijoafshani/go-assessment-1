package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartRestServer() error {
	gin.SetMode(gin.ReleaseMode)
	// TODO fix gin logger

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return err
	}

	return nil
}
