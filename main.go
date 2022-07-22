package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()
	setupRouters(router)
	router.Run(":8080")

}

func setupRouters(c *gin.Engine) {
	c.GET("/register",registerClient)
}
