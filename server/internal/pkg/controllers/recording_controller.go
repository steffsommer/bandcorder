package controllers

import (
	"net/http"
	"server/internal/pkg/interfaces"

	"github.com/gin-gonic/gin"
)

type RecordingController struct {
	recorder interfaces.Recorder
}

// NewRecordingController creates a new RecordingController
func NewRecordingController(
	recorder interfaces.Recorder,
) *RecordingController {
	return &RecordingController{
		recorder: recorder,
	}
}

// HandleStart starts a new recording
// Returns 200 OK on success, 500 Internal Server Error on failure
func (r RecordingController) HandleStart(c *gin.Context) {
	err := r.recorder.Start()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

// HandleStop stops the current recording
// Returns 200 OK on success, 500 Internal Server Error on failure
func (r RecordingController) HandleStop(c *gin.Context) {
	err := r.recorder.Stop()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

// HandleAbort aborts the current recording
// Returns 200 OK on success, 500 Internal Server Error on failure
func (r RecordingController) HandleAbort(c *gin.Context) {
	err := r.recorder.Abort()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
