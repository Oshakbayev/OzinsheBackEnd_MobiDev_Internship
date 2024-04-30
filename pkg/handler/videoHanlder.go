package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"ozinshe/pkg/entity"
	"ozinshe/pkg/helpers"
	"strconv"
)

func (h *Handler) AddNewSeries(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // Parse up to 10 MB of data
	if err != nil {
		h.log.Printf("error during ParseMultipartForm in CreateEvent(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid multipart form")
		return
	}
	formData := c.Request.MultipartForm
	movieIDstr := c.Param("id")
	movieID, err := strconv.Atoi(movieIDstr)

	if err != nil {
		h.log.Print("movieId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	VideoFileHeaders := formData.File["video[]"]
	var videosLinks []string
	videoDirectoryLink, err := h.svc.GetVideoDirectoryLinkByMovieId(movieID)
	if err != nil {
		h.log.Printf("error during AddNewSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add series")
		return
	}
	for _, file := range VideoFileHeaders {
		dst := entity.UploadedFilesPath + videoDirectoryLink + "/" + helpers.GenerateRandomKey(entity.UploadLinkNameLength)
		videosLinks = append(videosLinks, dst)
		// Upload the file to specific dst
		if err := c.SaveUploadedFile(file, dst); err != nil {
			h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	seasonIdstr := c.Param("seasonId")
	seasonId, err := strconv.Atoi(seasonIdstr)
	if err != nil {
		h.log.Print("seasonId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "seasonId is not a number")
		return
	}
	err = h.svc.AddSeries(movieID, seasonId, videosLinks)
	if err != nil {
		h.log.Printf("error during AddNewSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add series")
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "series added")
}

func (h *Handler) AddNewSeason(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // Parse up to 10 MB of data
	if err != nil {
		h.log.Printf("error during ParseMultipartForm in CreateEvent(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid multipart form")
		//http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}
	formData := c.Request.MultipartForm
	movieIDstr := c.Param("id")
	movieID, err := strconv.Atoi(movieIDstr)
	if err != nil {
		h.log.Print("movieId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	videoDirectoryLink, err := h.svc.GetVideoDirectoryLinkByMovieId(movieID)
	if err != nil {
		h.log.Printf("error during AddNewSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add series")
		return
	}
	VideoFileHeaders := formData.File["video[]"]
	var videosLinks []string
	for _, file := range VideoFileHeaders {
		dst := entity.UploadedFilesPath + videoDirectoryLink + "/" + helpers.GenerateRandomKey(entity.UploadLinkNameLength)
		videosLinks = append(videosLinks, dst)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	err = h.svc.AddSeason(movieID, videosLinks)
	if err != nil {
		h.log.Printf("error during AddNewSeason(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add season")
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "season added")

}

func (h *Handler) DeleteMovieSeason(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Print("movieId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	seasonID, err := strconv.Atoi(c.Param("seasonId"))
	if err != nil {
		h.log.Print("seasonId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "seasonId is not a number")
		return
	}
	err = h.svc.DeleteMovieSeason(movieID, seasonID)
	if err != nil {
		h.log.Printf("error during DeleteMovieSeason(handler): %v", err)
		if os.IsNotExist(err) {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "File does not exist")
		} else if os.IsPermission(err) {
			h.WriteHTTPResponse(c, http.StatusForbidden, "Permission denied")
		}
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to delete season")
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "season deleted")
}

func (h *Handler) DeleteMovieSeries(c *gin.Context) {
	movieIDstr := c.Param("id")
	movieID, err := strconv.Atoi(movieIDstr)
	if err != nil {
		h.log.Printf("error during DeleteMovieSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	seasonIdstr := c.Param("seasonId")
	seasonId, err := strconv.Atoi(seasonIdstr)
	if err != nil {
		h.log.Printf("error during DeleteMovieSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "seasonId is not a number")
		return
	}
	seriesIdstr := c.Param("seriesId")
	seriesId, err := strconv.Atoi(seriesIdstr)
	if err != nil {
		h.log.Printf("error during DeleteMovieSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "seriesId is not a number")
		return
	}
	err = h.svc.DeleteMovieSeries(movieID, seasonId, seriesId)
	if err != nil {
		h.log.Printf("error during DeleteMovieSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to delete series")
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "series deleted")
}

func (h *Handler) UpdateSeason(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // Parse up to 10 MB of data
	if err != nil {
		h.log.Printf("error during ParseMultipartForm in UpdateSeason(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid multipart form")
		//http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}
	formData := c.Request.MultipartForm
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Print("movieId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	seasonID, err := strconv.Atoi(c.Param("seasonId"))
	if err != nil {
		h.log.Print("seasonId is not a number")
		h.WriteHTTPResponse(c, http.StatusBadRequest, "seasonId is not a number")
		return
	}
	videoDirectoryLink, err := h.svc.GetVideoDirectoryLinkByMovieId(movieID)
	if err != nil {
		h.log.Printf("error during UpdateSeason(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add series")
		return
	}
	err = h.svc.DeleteMovieSeason(movieID, seasonID)
	if err != nil {
		h.log.Printf("error during UpdateSeason(handler): %v", err)
		if os.IsNotExist(err) {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "File does not exist")
		} else if os.IsPermission(err) {
			h.WriteHTTPResponse(c, http.StatusForbidden, "Permission denied")
		}
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to delete season")
		return
	}
	VideoFileHeaders := formData.File["video[]"]
	var videosLinks []string
	for _, file := range VideoFileHeaders {
		dst := entity.UploadedFilesPath + videoDirectoryLink + "/" + helpers.GenerateRandomKey(entity.UploadLinkNameLength)
		videosLinks = append(videosLinks, dst)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	err = h.svc.AddSeason(movieID, videosLinks)
	if err != nil {
		h.log.Printf("error during UpdateSeason(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add season")
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "season updated")
}

func (h *Handler) UpdateSeries(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // Parse up to 10 MB of data
	if err != nil {
		h.log.Printf("error during ParseMultipartForm in UpdateSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid multipart form")
		return
	}
	formData := c.Request.MultipartForm
	movieIDstr := c.Param("id")
	movieID, err := strconv.Atoi(movieIDstr)
	if err != nil {
		h.log.Printf("error during UpdateSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "movieId is not a number")
		return
	}
	seasonIdstr := c.Param("seasonId")
	seasonId, err := strconv.Atoi(seasonIdstr)
	if err != nil {
		h.log.Printf("error during UpdateSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "seasonId is not a number")
		return
	}
	seriesIdstr := c.Param("seriesId")
	seriesId, err := strconv.Atoi(seriesIdstr)
	if err != nil {
		h.log.Printf("error during UpdateSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "seriesId is not a number")
		return
	}
	videoDirectoryLink, err := h.svc.GetVideoDirectoryLinkByMovieId(movieID)
	if err != nil {
		h.log.Printf("error during UpdateSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add series")
		return
	}
	VideoFile, ok := formData.File["video"]
	if !ok {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, "Something wrong with VideoFile")
		return
	}
	dst := videoDirectoryLink + "/" + helpers.GenerateRandomKey(entity.UploadLinkNameLength)
	if err := c.SaveUploadedFile(VideoFile[0], dst); err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.svc.UpdateSeries(movieID, seasonId, seriesId, dst)
	if err != nil {
		h.log.Printf("error during UpdateSeries(handler): %v", err)
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Failed to add series")
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "series updated")
}
