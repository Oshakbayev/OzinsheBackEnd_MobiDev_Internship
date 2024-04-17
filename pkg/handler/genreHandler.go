package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllGenres(c *gin.Context) {
	genres, err := h.svc.GetAllGenres()
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, genres)
}

//func (h *Handler) GetMovieMainsByGenres(c *gin.Context) {
//	categoryId, err := strconv.Atoi(c.Param("genreId"))
//	if err != nil {
//		h.log.Print("genreId is not a number")
//		h.WriteHTTPResponse(c, http.StatusBadRequest, "genreId is not a number")
//		return
//	}
//	params := c.Request.URL.Query()
//	limit, err := strconv.Atoi(params.Get("limit"))
//	if err != nil {
//		h.log.Print("limit is not a number")
//		h.WriteHTTPResponse(c, http.StatusBadRequest, "limit is not a number")
//		return
//	}
//	offset, err := strconv.Atoi(params.Get("offset"))
//	if err != nil {
//		h.log.Print("offset is not a number")
//		h.WriteHTTPResponse(c, http.StatusBadRequest, "offset is not a number")
//		return
//	}
//	movieMains, err := h.svc.GetMovieMainsBy(categoryId, limit, offset)
//	if err != nil {
//		h.log.Print("error in GetMovieMainsByCategory(handler)")
//		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, movieMains)
//}
