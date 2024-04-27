package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ozinshe/pkg/entity"
)

func (h *Handler) CreateFavoriteMovie(c *gin.Context) {
	var favorite entity.Favorite
	if err := c.BindJSON(&favorite); err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	favorite.UserId = c.Value("decodedClaims").(*entity.Claims).Sub
	if err := h.svc.CreatFavoriteMovie(favorite); err != nil {
		if err.Error() == "409" {
			h.WriteHTTPResponse(c, http.StatusConflict, "Invalid input body")
			return
		}
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(c, http.StatusCreated, "movie added")
}

//func (h *Handler) GetFavoriteMovies(c *gin.Context) {
//	userId := c.Value("decodedClaims").(*entity.Claims).Sub
//	movieIds, err := h.svc.GetUserFavoriteMovieIDs(userId)
//	if err != nil && err.Error() != entity.ErrNoRows {
//		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	response := map[string]interface{}{
//		"movieIds": movieIds,
//	}
//
//	c.JSON(http.StatusOK, response)
//}

func (h *Handler) DeleteFavoriteMovies(c *gin.Context) {
	var favorite entity.Favorite
	if err := c.BindJSON(&favorite); err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	favorite.UserId = c.Value("decodedClaims").(*entity.Claims).Sub
	err := h.svc.DeleteFavoriteMovie(favorite)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "movie deleted")
}
