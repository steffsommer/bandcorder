package facades

import (
	"server/internal/pkg/interfaces"

	"github.com/sirupsen/logrus"
)

type RecordingFacade struct {
	eventbus interfaces.RecordingEventBus
	recorder interfaces.Recorder
}

func NewRecordingFacade(
	eventbus interfaces.RecordingEventBus,
	recorder interfaces.Recorder,
) *RecordingFacade {
	return &RecordingFacade{
		eventbus: eventbus,
		recorder: recorder,
	}
}

func (r *RecordingFacade) Start() error {
	err := r.recorder.Start()
	if err != nil {
		logrus.Errorf("Failed to start recording: %s", err.Error())
		return err
	}
	r.eventbus.NotifyStarted()
	logrus.Info("Recording started successfully")
	return err
}

func (r *RecordingFacade) Stop() error {
	err := r.recorder.Stop()
	if err != nil {
		logrus.Errorf("Failed to start recording: %s", err.Error())
		return err
	}
	r.eventbus.NotifyStarted()
	logrus.Info("Recording started successfully")
	return err
}

func (r *RecordingFacade) Abort() error {
	err := r.recorder.Abort()
	if err != nil {
		logrus.Errorf("Failed to start recording: %s", err.Error())
		return err
	}
	r.eventbus.NotifyStarted()
	logrus.Info("Recording started successfully")
	return err
}
