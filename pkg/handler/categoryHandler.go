package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ozinshe/pkg/entity"
	"strconv"
)

func (h *Handler) GetMovieMainsByAllCategory(c *gin.Context) {
	categoryIds, err := h.svc.GetAllCategories()
	if err != nil {
		h.log.Print("genreId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "genreId is not a number")
		return
	}
	userId := c.Value("decodedClaims").(*entity.Claims).Sub
	responseMap := make(map[string]struct {
		MovieMains []entity.MovieMain
		CategoryId int
	})
	for _, category := range categoryIds {
		movieMains, err := h.svc.GetAllMovieMainsByCategory(userId, category.Id)
		if err != nil {
			h.log.Println("error in GetMovieMainsByAllCategory(handler):", err)
			h.WriteHTTPResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		responseMap[category.Name] = struct {
			MovieMains []entity.MovieMain
			CategoryId int
		}{MovieMains: movieMains, CategoryId: category.Id}
	}
	c.JSON(http.StatusOK, responseMap)
}

func (h *Handler) GetMovieMainsByCategory(c *gin.Context) {
	params := c.Request.URL.Query()
	userId := c.Value("decodedClaims").(*entity.Claims).Sub
	categoryId, err := strconv.Atoi(params.Get("categoryId"))
	if err != nil {
		h.log.Print("limit is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "limit is not a number")
		return
	}
	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		h.log.Print("limit is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "limit is not a number")
		return
	}
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		h.log.Print("offset is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "offset is not a number")
		return
	}
	movieMains, err := h.svc.GetMovieMainsByCategory(userId, categoryId, limit, offset)
	if err != nil {
		h.log.Println("error in GetMovieMainsByCategory(handler)", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, movieMains)
}
