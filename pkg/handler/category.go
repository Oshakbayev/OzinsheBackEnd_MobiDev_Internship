package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := h.svc.GetAllCategories()
	if err != nil {
		h.log.Print("error in GetCategories(handler)")
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, categories)
}
