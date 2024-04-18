package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ozinshe/pkg/entity"
	"strconv"
)

func (h *Handler) GetMovieMainsByCategory(c *gin.Context) {
	var movieMains []entity.MovieMain
	params := c.Request.URL.Query()
	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		h.log.Print("genreId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "genreId is not a number")
		return
	}

	if params.Get("limit") == "" || params.Get("offset") == "" {
		movieMains, err = h.svc.GetAllMovieMainsByCategory(categoryId)
	} else {
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
		movieMains, err = h.svc.GetMovieMainsByCategory(categoryId, limit, offset)

	}
	if err != nil {
		h.log.Print("error in GetMovieMainsByCategory(handler)")
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movieMains)
}
