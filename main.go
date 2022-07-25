package main

import "github.com/gin-gonic/gin"

type Quotes struct {
	Q string `json:"q"`
	A string `json:"a"`
	H string `json:"h"`
}

func main() {

	// zenquotes()

	router := gin.Default()
	setupRouters(router)
	router.Run(":8080")

}

func setupRouters(c *gin.Engine) {
	c.GET("/register", registerClient)
}
