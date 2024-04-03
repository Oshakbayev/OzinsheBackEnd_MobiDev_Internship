package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ozinshe/pkg/entity"
	"strconv"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.svc.SignUp(&user)
	if err != nil {
		if err.Error() == entity.AlreadyExist.Error() {
			h.WriteHTTPResponse(c, http.StatusBadRequest, entity.AlreadyExist.Error())
			return
		} else {
			h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	h.WriteHTTPResponse(c, http.StatusOK, "user registered with id "+strconv.Itoa(user.Id))

}

func (h *Handler) VerifyAccount(c *gin.Context) {
	secretCode := c.Query("link")
	err := h.svc.VerifyAccount(secretCode)
	if err != nil {
		if err.Error() == entity.DidNotFind.Error() {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "invalid  verification link")
			return
		} else if err.Error() == entity.ExpiredLink.Error() {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "verification link expired")
			return
		}
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
	}
	h.WriteHTTPResponse(c, http.StatusOK, "account confirmed!")
}

func (h *Handler) HomePageHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"Message": "Hello,World!",
	})
}
