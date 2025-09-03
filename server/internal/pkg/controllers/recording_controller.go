package controllers

import (
	"net/http"
	"server/internal/pkg/interfaces"

	"github.com/gin-gonic/gin"
)

type RecordingController struct {
	recorder interfaces.Recorder
	notifier interfaces.RecordingStateNotifier
}

func NewRecordingController(
	recorder interfaces.Recorder,
	notifier interfaces.RecordingStateNotifier,
) RecordingController {
	return RecordingController{
		recorder: recorder,
		notifier: notifier,
	}
}

// HandleStart starts a new recording
func (r RecordingController) HandleStart(c *gin.Context) {
	err := r.recorder.Start()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

// HandleStop stops the current recording
func (r RecordingController) HandleStop(c *gin.Context) {
	err := r.recorder.Stop()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

// HandleAbort stops the current recording
func (r RecordingController) HandleAbort(c *gin.Context) {
	err := r.recorder.Abort()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
