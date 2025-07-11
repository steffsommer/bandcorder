package main

import (
	recorder "steffsommer/bandcorder/internal/pkg"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting bandcorder")
	recorder, err := recorder.NewRecorder()
	if err != nil {
		logrus.Errorf("Failed to create recorder instance: %s", err.Error())
	}
	if err := recorder.Start(); err != nil {
		logrus.Errorf("Failed to start recording: %s", err.Error())
	}
	time.Sleep(5 * time.Second)
	if err := recorder.Stop(); err != nil {
		logrus.Errorf("Failed to stop recording: %s", err.Error())
	}

}
