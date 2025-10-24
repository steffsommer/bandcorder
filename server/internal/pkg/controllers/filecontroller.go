package controllers

import (
	"encoding/json"
	"net/http"
	"server/internal/pkg/interfaces"

	"github.com/gin-gonic/gin"
)

type RenameDto struct {
	FileName string `validate:"required,regex=^[a-zA-Z0-9._-]+$"`
}

type FileController struct {
	fileStorage interfaces.StorageService
}

func NewFileController(fileStorage interfaces.StorageService) *FileController {
	return &FileController{
		fileStorage: fileStorage,
	}
}

func (f *FileController) HandleRenameLast(c *gin.Context) {
	var dto RenameDto
	err := json.NewDecoder(c.Request.Body).Decode(&dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = f.fileStorage.RenameLastRecording(dto.FileName)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
