package v1

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	l "api-gateway/pkg/logger"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}


// File upload
// @Security    BearerAuth
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Router /v1/imageupload [post]  // Changed the endpoint to /v1/imageupload
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
func (h *handlerV1) UploadImageFile(c *gin.Context) {
	var file File

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Error while uploading file", l.Any("post", err))
		return
	}

	// Check if the file has a valid image file extension
	allowedExtensions := []string{".png", ".jpg", ".jpeg"}
	validExtension := false
	for _, ext := range allowedExtensions {
		if filepath.Ext(file.File.Filename) == ext {
			validExtension = true
			break
		}
	}

	if !validExtension {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't find matching image file format",
		})
		h.log.Error("Error while uploading image file", l.Any("image-upload", err))
		return
	}

	id := uuid.New()
	fileName := id.String() + filepath.Ext(file.File.Filename)
	dst, _ := os.Getwd()

	// Update the directory path to "media/image"
	if _, err := os.Stat(dst + "/media/image"); os.IsNotExist(err) {
		os.Mkdir(dst+"/media/image", os.ModePerm)
	}

	filePath := "/media/image/" + fileName
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
