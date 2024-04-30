package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ozinshe/pkg/entity"
)

func (h *Handler) GetUserProfile(c *gin.Context) {
	claims := c.Value("decodedClaims").(*entity.Claims)
	userProfile, err := h.svc.GetUserProfileByUserId(claims.Sub)
	if err != nil {
		if err.Error() == entity.ErrNoRows {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "this user have no information")
		}
		h.log.Print("error in GetUserProfile(handler)")
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userProfile)
}

func (h *Handler) UpdateUserProfile(c *gin.Context) {
	userProfile := entity.UserProfile{}
	if err := c.BindJSON(&userProfile); err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body"+err.Error())
		return
	}
	claims := c.Value("decodedClaims").(*entity.Claims)
	userProfile.UserId = claims.Sub
	err := h.svc.UpdateUserProfile(&userProfile)
	if err != nil {
		h.log.Print("error in UserProfile(handler) during create userProfile")
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "user information updated")
}
