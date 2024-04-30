package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ozinshe/pkg/entity"
	"ozinshe/pkg/helpers"
	"strconv"
	"strings"
)

func (h *Handler) CreateMovie(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // Parse up to 10 MB of data
	if err != nil {
		h.log.Printf("error during ParseMultipartForm in CreateEvent(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid multipart form")
		//http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}
	formData := c.Request.MultipartForm

	movie := entity.Movie{}
	if err := json.NewDecoder(strings.NewReader(formData.Value["json"][0])).Decode(&movie); err != nil {
		log.Println(err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	ScreenshotFileHeaders := formData.File["screenshots[]"]
	movie.VideoDirectoryLink = entity.UploadedFilesPath + helpers.GenerateRandomKey(entity.UploadLinkNameLength)
	PosterFile, ok := formData.File["poster"]
	if !ok {
		log.Println(PosterFile)
		h.WriteHTTPResponse(c, http.StatusInternalServerError, "Something wrong with PosterFile")
		return
	}
	VideoFileHeaders := formData.File["video[]"]
	movie.PosterLink = entity.UploadedFilesPath + helpers.GenerateRandomKey(entity.UploadLinkNameLength)
	if err := c.SaveUploadedFile(PosterFile[0], movie.PosterLink); err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for _, file := range ScreenshotFileHeaders {
		dst := entity.UploadedFilesPath + helpers.GenerateRandomKey(entity.UploadLinkNameLength)
		movie.ScreenshotLinks = append(movie.ScreenshotLinks, dst)
		// Upload the file to specific dst
		if err := c.SaveUploadedFile(file, dst); err != nil {
			h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	for _, file := range VideoFileHeaders {
		dst := movie.VideoDirectoryLink + "/" + helpers.GenerateRandomKey(entity.UploadLinkNameLength)
		movie.VideoLinks = append(movie.VideoLinks, dst)
		// Upload the file to specific dst
		if err := c.SaveUploadedFile(file, dst); err != nil {
			h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if err := h.svc.CreateMovie(&movie, formData); err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(c, http.StatusCreated, "movie created!")
}

func (h *Handler) GetAllMovies(c *gin.Context) {
	params := c.Request.URL.Query()
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
	userId := c.Value("decodedClaims").(*entity.Claims).Sub
	movies, err := h.svc.GetAllMovies(userId, limit, offset)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (h *Handler) GetMovieById(c *gin.Context) {
	params := c.Param("id")
	movieID, err := strconv.Atoi(params)
	if err != nil {
		h.log.Print("movieId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	userId := c.Value("decodedClaims").(*entity.Claims).Sub
	movie, err := h.svc.GetMovieById(userId, movieID)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (h *Handler) UpdateMovieById(c *gin.Context) {
	params := c.Param("id")
	movieID, err := strconv.Atoi(params)
	if err != nil {
		h.log.Print("movieId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	movie := entity.Movie{}
	if err := c.BindJSON(&movie); err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	movie.Id = movieID
	err = h.svc.UpdateMovieById(&movie)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "movie updated!")
}

func (h *Handler) DeleteMovieById(c *gin.Context) {
	params := c.Param("id")
	movieID, err := strconv.Atoi(params)
	if err != nil {
		h.log.Print("movieId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	err = h.svc.DeleteMovieById(movieID)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "movie deleted!")
}

func (h *Handler) GetMovieSeasonById(c *gin.Context) {
	seasonId := c.Param("seasonId")
	movieId := c.Param("id")
	intSeasonId, err := strconv.Atoi(seasonId)
	if err != nil {
		h.log.Print("seasonId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "seasonId is not a number")
		return
	}
	intMovieId, err := strconv.Atoi(movieId)
	if err != nil {
		h.log.Print("seasonId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	seasonSeriesLinks, err := h.svc.GetMovieSeason(intMovieId, intSeasonId)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, seasonSeriesLinks)
}

func (h *Handler) GetMovieSeriesById(c *gin.Context) {
	seasonId := c.Param("seasonId")
	movieId := c.Param("movieId")
	episodeId := c.Param("seriesId")
	intSeasonId, err := strconv.Atoi(seasonId)
	if err != nil {
		h.log.Print("seasonId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "seasonId is not a number")
		return
	}
	intMovieId, err := strconv.Atoi(movieId)
	if err != nil {
		h.log.Print("seasonId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	intEpisodeId, err := strconv.Atoi(episodeId)
	if err != nil {
		h.log.Print("seasonId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "episodeId is not a number")
		return
	}
	seasonSeriesLinks, err := h.svc.GetMovieSeries(intMovieId, intEpisodeId, intSeasonId)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, seasonSeriesLinks)
}

func (h *Handler) GetMovieMainsByTitle(c *gin.Context) {
	params := c.Request.URL.Query()
	title := params.Get("title")
	movieMains, err := h.svc.GetMovieMainsByTitle(title)
	if err != nil {
		h.log.Print("error in GetMovieMainsByTitle(handler)")
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movieMains)
}
