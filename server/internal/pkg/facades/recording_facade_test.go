package facades

import (
	"errors"
	"testing"
	"time"

	"server/internal/pkg/interfaces"
	"server/internal/pkg/testutils/mocks"

	"github.com/stretchr/testify/assert"
)

func TestRecordingFacade_Start_Success(t *testing.T) {
	rec := mocks.NewMockRecorder(t)
	bus := mocks.NewMockRecordingEventBus(t)
	meta := getTestRecordingMetaData()
	rec.EXPECT().Start().Return(meta, nil)
	bus.EXPECT().NotifyStarted(meta).Return()
	facade := NewRecordingFacade(bus, rec)

	res, err := facade.Start()

	assert.NoError(t, err)
	assert.Equal(t, meta, res)
}

func TestRecordingFacade_Start_Error(t *testing.T) {
	rec := mocks.NewMockRecorder(t)
	bus := mocks.NewMockRecordingEventBus(t)
	meta := getTestRecordingMetaData()
	rec.EXPECT().Start().Return(meta, errors.New("start failed"))
	facade := NewRecordingFacade(bus, rec)

	res, err := facade.Start()

	assert.Error(t, err)
	assert.Equal(t, meta, res)
}

func TestRecordingFacade_Stop_Success(t *testing.T) {
	rec := mocks.NewMockRecorder(t)
	bus := mocks.NewMockRecordingEventBus(t)
	rec.EXPECT().Stop().Return(nil)
	bus.EXPECT().NotifyStopped().Return()
	facade := NewRecordingFacade(bus, rec)

	err := facade.Stop()

	assert.NoError(t, err)
}

func TestRecordingFacade_Stop_Error(t *testing.T) {
	rec := mocks.NewMockRecorder(t)
	bus := mocks.NewMockRecordingEventBus(t)
	rec.EXPECT().Stop().Return(errors.New("stop failed"))
	facade := NewRecordingFacade(bus, rec)

	err := facade.Stop()

	assert.Error(t, err)
}

func TestRecordingFacade_Abort_Success(t *testing.T) {
	rec := mocks.NewMockRecorder(t)
	bus := mocks.NewMockRecordingEventBus(t)
	rec.EXPECT().Abort().Return(nil)
	bus.EXPECT().NotifyStopped().Return()
	facade := NewRecordingFacade(bus, rec)

	err := facade.Abort()

	assert.NoError(t, err)
}

func TestRecordingFacade_Abort_Error(t *testing.T) {
	rec := mocks.NewMockRecorder(t)
	bus := mocks.NewMockRecordingEventBus(t)
	rec.EXPECT().Abort().Return(errors.New("fail"))
	facade := NewRecordingFacade(bus, rec)

	err := facade.Abort()

	assert.Error(t, err)
}

func TestRecordingFacade_GetMic(t *testing.T) {
	rec := mocks.NewMockRecorder(t)
	bus := mocks.NewMockRecordingEventBus(t)
	rec.EXPECT().GetMic().Return("mic1", nil)
	facade := NewRecordingFacade(bus, rec)

	mic, err := facade.GetMic()

	assert.NoError(t, err)
	assert.Equal(t, "mic1", mic)
}

func getTestRecordingMetaData() interfaces.RecordingMetaData {
	return interfaces.RecordingMetaData{
		FileName: "test-file",
		Started:  time.Now(),
	}
}
