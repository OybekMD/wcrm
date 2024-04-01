package v1

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	l "api-gateway/pkg/logger"
)

// File upload
// @Security    BearerAuth
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Router /v1/pdfupload [post]  // Changed the endpoint to /v1/pdfupload
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
func (h *handlerV1) UploadPDFFile(c *gin.Context) {
	var file File

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Error while uploading file", l.Any("post", err))
		return
	}

	// Check if the file has a valid PDF file extension
	allowedExtensions := []string{".pdf"}
	validExtension := false
	for _, ext := range allowedExtensions {
		if filepath.Ext(file.File.Filename) == ext {
			validExtension = true
			break
		}
	}

	if !validExtension {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't find matching PDF file format",
		})
		h.log.Error("Error while uploading PDF file", l.Any("pdf-upload", err))
		return
	}

	dst, _ := os.Getwd()

	// Update the directory path to "media/pdf"
	if _, err := os.Stat(dst + "/media/pdf"); os.IsNotExist(err) {
		os.Mkdir(dst+"/media/pdf", os.ModePerm)
	}

	// Replace spaces with underscores in the file name
	fileName := strings.ReplaceAll(file.File.Filename, " ", "_")

	filePath := "/media/pdf/" + fileName
	err = c.SaveUploadedFile(file.File, dst+filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Error while uploading PDF file",
		})
		h.log.Error("Error while uploading PDF file", l.Any("post", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": c.Request.Host + filePath,
	})
}
