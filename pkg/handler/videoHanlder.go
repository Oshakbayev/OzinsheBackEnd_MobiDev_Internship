package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) AddNewSeason(c *gin.Context) {
	formData := c.Request.MultipartForm
	movieIDstr := c.Param("movieId")
	movieName := c.Param("movieName")
	movieID, err := strconv.Atoi(movieIDstr)

	if err != nil {
		h.log.Print("movieId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	VideoFileHeaders := formData.File["video[]"]
	var videosLinks []string
	for _, file := range VideoFileHeaders {
		log.Println(file.Filename)
		dst := "/assets/uploads/" + movieName + "/videos/" + file.Filename
		videosLinks = append(videosLinks, dst)
		// Upload the file to specific dst
		if err := c.SaveUploadedFile(file, dst); err != nil {
			h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	var reqBody map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(formData.Value["json"][0])).Decode(&reqBody); err != nil {
		log.Println(err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	seasonIdstr := c.Request.URL.Query().Get("seasonId")
	if seasonIdstr != "" {
		seasonId, err := strconv.Atoi(seasonIdstr)
		if err != nil {
			h.log.Print("seasonId is not a number")
			h.WriteHTTPResponse(c, http.StatusBadRequest, "seasonId is not a number")
			return
		}
		err = h.svc.AddSeries(movieID, seasonId, videosLinks)
	} else {
		err = h.svc.AddSeason(movieID, videosLinks)
	}
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add season")
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "series added")
}

func (h *Handler) AddNewSeries(c *gin.Context) {
	formData := c.Request.MultipartForm
	movieIDstr := c.Param("movieId")
	seasonIdstr := c.Param("seasonId")
	seasonId, err := strconv.Atoi(seasonIdstr)
	movieName := c.Param("movieName")
	movieID, err := strconv.Atoi(movieIDstr)
	if err != nil {
		h.log.Print("movieId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	VideoFileHeaders := formData.File["video[]"]
	var videosLinks []string
	for _, file := range VideoFileHeaders {

		log.Println(file.Filename)
		dst := "/assets/uploads/" + movieName + "/videos/" + file.Filename
		videosLinks = append(videosLinks, dst)
		// Upload the file to specific dst
		if err := c.SaveUploadedFile(file, dst); err != nil {
			h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	var reqBody map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(formData.Value["json"][0])).Decode(&reqBody); err != nil {
		log.Println(err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err = h.svc.AddSeries(movieID, seasonId, videosLinks)
	h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add season")
	return
}
