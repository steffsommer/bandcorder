package controllers

import (
	"encoding/json"
	"net/http"
	"server/internal/pkg/interfaces"

	"github.com/gin-gonic/gin"
)

type UpdateBpmDto struct {
	Bpm uint8
}

type OnOffStateDto struct {
	On bool
}

type MetronomeController struct {
	metronome interfaces.Metronome
}

func NewMetronomeController(metronome interfaces.Metronome) *MetronomeController {
	return &MetronomeController{
		metronome: metronome,
	}
}

func (m *MetronomeController) HandleSwitchOnState(c *gin.Context) {
	var dto OnOffStateDto
	err := json.NewDecoder(c.Request.Body).Decode(&dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if dto.On {
		err = m.metronome.Start()
	} else {
		err = m.metronome.Stop()
	}
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (m *MetronomeController) HandleBpmUpdate(c *gin.Context) {
	var dto UpdateBpmDto
	err := json.NewDecoder(c.Request.Body).Decode(&dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = m.metronome.UpdateBpm(int(dto.Bpm)); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
