package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) AddNewSeason(c *gin.Context) {
	formData := c.Request.MultipartForm
	VideoFileHeaders := formData.File["video[]"]
	var reqbody map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(formData.Value["json"][0])).Decode(&reqbody); err != nil {
		log.Println(err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.svc.AddSeason()
}
