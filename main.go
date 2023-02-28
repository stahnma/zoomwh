package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type zoomWebhook struct {
}

func dostuff(c *gin.Context) {

	c.JSON(http.StatusOK, "Did some stuff")

}

func main() {

	router := gin.Default()
	router.POST("/zoom", dostuff)

	router.Run("localhost:9999")
}
