package api

import (
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	//TODO impl
}

func getBalance(c *gin.Context) {
	//TODO impl
}

func getAllBalances(c *gin.Context) {
	//TODO impl
}

func addBalance(c *gin.Context) {
	//TODO impl
}

func addToAllBalances(c *gin.Context) {
	//TODO impl
}

func StartRestServer() error {
	gin.SetMode(gin.ReleaseMode)
	// TODO fix gin logger

	r := gin.New()
	r.POST("/create", create)
	r.GET("/getBalance", getBalance)
	r.GET("/getAllBalances", getAllBalances)
	r.PUT("/addBalance", addBalance)
	r.PUT("/addToAllBalances", addToAllBalances)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return err
	}

	return nil
}
