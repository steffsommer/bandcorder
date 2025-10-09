package facades

import (
	"server/internal/pkg/interfaces"

	"github.com/sirupsen/logrus"
)

// RecordingFacade is a proxy implementation of the recorder interface, which notifies
// the event system about recording state changes
type RecordingFacade struct {
	eventbus interfaces.RecordingEventBus
	recorder interfaces.Recorder
}

// NewRecordingFacade creates a new NewRecordingFacade
func NewRecordingFacade(
	eventbus interfaces.RecordingEventBus,
	recorder interfaces.Recorder,
) *RecordingFacade {
	return &RecordingFacade{
		eventbus: eventbus,
		recorder: recorder,
	}
}

// Start starts a new recording and notifies the event bus
func (r *RecordingFacade) Start() (interfaces.RecordingMetaData, error) {
	res, err := r.recorder.Start()
	if err != nil {
		logrus.Errorf("Failed to start recording: %s", err.Error())
		return res, err
	}
	r.eventbus.NotifyStarted(res)
	logrus.Info("Recording started successfully")
	return res, nil
}

// Stop stops the current recording and notifies the event bus
func (r *RecordingFacade) Stop() error {
	err := r.recorder.Stop()
	if err != nil {
		logrus.Errorf("Failed to stop recording: %s", err.Error())
		return err
	}
	r.eventbus.NotifyStopped()
	logrus.Info("Recording stopped successfully")
	return err
}

// Abort aborts the current recording and notifies the event bus
func (r *RecordingFacade) Abort() error {
	err := r.recorder.Abort()
	if err != nil {
		logrus.Errorf("Failed to abort recording: %s", err.Error())
		return err
	}
	r.eventbus.NotifyStopped()
	logrus.Info("Recording aborted successfully")
	return err
}

// GetMic returns the name of the microphone in use by the recorder
func (r *RecordingFacade) GetMic() (string, error) {
	return r.recorder.GetMic()
}
