package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HomePageHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"Message": "Hello,World!",
	})
}
