package facades

import (
	"errors"
	"testing"
	"time"

	"server/internal/pkg/interfaces"
	"server/internal/pkg/testutils/mocks"

	"github.com/stretchr/testify/assert"
)

var (
	recorder        *mocks.MockRecorder
	eventBus        *mocks.MockRecordingEventBus
	playbackService *mocks.MockPlaybackService
)

func setup(t *testing.T) *RecordingFacade {
	recorder = mocks.NewMockRecorder(t)
	eventBus = mocks.NewMockRecordingEventBus(t)
	playbackService = mocks.NewMockPlaybackService(t)
	return NewRecordingFacade(eventBus, recorder, playbackService)
}

func TestRecordingFacade_Start_Success(t *testing.T) {
	facade := setup(t)
	meta := getTestRecordingMetaData()
	recorder.EXPECT().Start().Return(meta, nil)
	eventBus.EXPECT().NotifyStarted(meta).Return()
	playbackService.EXPECT().Play(interfaces.SwitchOn)

	res, err := facade.Start()

	assert.NoError(t, err)
	assert.Equal(t, meta, res)
}

func TestRecordingFacade_Start_Error(t *testing.T) {
	meta := getTestRecordingMetaData()
	facade := setup(t)
	recorder.EXPECT().Start().Return(meta, errors.New("start failed"))

	res, err := facade.Start()

	assert.Error(t, err)
	assert.Equal(t, meta, res)
}

func TestRecordingFacade_Stop_Success(t *testing.T) {
	facade := setup(t)
	recorder.EXPECT().Stop().Return(nil)
	eventBus.EXPECT().NotifyStopped().Return()
	playbackService.EXPECT().Play(interfaces.SwitchOff)

	err := facade.Stop()

	assert.NoError(t, err)
}

func TestRecordingFacade_Stop_Error(t *testing.T) {
	facade := setup(t)
	recorder.EXPECT().Stop().Return(errors.New("stop failed"))

	err := facade.Stop()

	assert.Error(t, err)
}

func TestRecordingFacade_Abort_Success(t *testing.T) {
	facade := setup(t)
	recorder.EXPECT().Abort().Return(nil)
	eventBus.EXPECT().NotifyStopped().Return()
	playbackService.EXPECT().Play(interfaces.Delete)

	err := facade.Abort()

	assert.NoError(t, err)
}

func TestRecordingFacade_Abort_Error(t *testing.T) {
	facade := setup(t)
	recorder.EXPECT().Abort().Return(errors.New("fail"))

	err := facade.Abort()

	assert.Error(t, err)
}

func TestRecordingFacade_GetMic(t *testing.T) {
	facade := setup(t)
	recorder.EXPECT().GetMic().Return("mic1", nil)

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
