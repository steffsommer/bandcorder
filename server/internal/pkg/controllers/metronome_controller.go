package controllers

import (
	"encoding/json"
	"net/http"
	"server/internal/pkg/interfaces"

	"github.com/gin-gonic/gin"
)

type UpdateBpmDto struct {
	Bpm uint8 `json:"bpm"`
}

type MetronomeController struct {
	metronome interfaces.Metronome
}

func NewMetronomeController(metronome interfaces.Metronome) *MetronomeController {
	return &MetronomeController{
		metronome: metronome,
	}
}

func (m *MetronomeController) HandleStart(c *gin.Context) {
	if err := m.metronome.Start(); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (m *MetronomeController) HandleStop(c *gin.Context) {
	if err := m.metronome.Stop(); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (m *MetronomeController) HandleGetState(c *gin.Context) {
	state := m.metronome.GetState()
	c.JSON(http.StatusOK, &state)
}

func (m *MetronomeController) HandleUpdateBpm(c *gin.Context) {
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
