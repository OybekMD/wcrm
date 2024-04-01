package v1

import (
	"net/http"
	"os"
	"path/filepath"

	l "api-gateway/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// File upload
// @Security    BearerAuth
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Router /v1/videoupload [post]  // Changed the endpoint to /v1/videoupload
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
func (h *handlerV1) UploadVideoFile(c *gin.Context) {
	var file File

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Error while uploading file", l.Any("post", err))
		return
	}

	// Check if the file has a valid video file extension
	allowedExtensions := []string{".mp4", ".avi", ".mkv"}
	validExtension := false
	for _, ext := range allowedExtensions {
		if filepath.Ext(file.File.Filename) == ext {
			validExtension = true
			break
		}
	}

	if !validExtension {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't find matching video file format",
		})
		h.log.Error("Error while uploading video file", l.Any("video-upload", err))
		return
	}

	id := uuid.New()
	fileName := id.String() + filepath.Ext(file.File.Filename)
	dst, _ := os.Getwd()

	// Update the directory path to "media/video"
	if _, err := os.Stat(dst + "/media/video"); os.IsNotExist(err) {
		os.Mkdir(dst+"/media/video", os.ModePerm)
	}

	filePath := "/media/video/" + fileName
	err = c.SaveUploadedFile(file.File, dst+filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Couldn't find matching information, Have you registered before?",
		})
		h.log.Error("Error while getting customer by email", l.Any("post", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": c.Request.Host + filePath,
	})
}
