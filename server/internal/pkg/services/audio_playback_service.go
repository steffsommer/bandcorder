package services

import (
	"bytes"
	"encoding/binary"
	"server/internal/pkg/interfaces"
	"server/resources"
	"sync"

	"github.com/gen2brain/malgo"
	"github.com/go-audio/wav"
	"github.com/sirupsen/logrus"
)

type AudioPlaybackService struct {
	fileMap map[interfaces.AudioEffect]string
	ctx     *malgo.AllocatedContext
}

func NewAudioPlaybackService() *AudioPlaybackService {
	return &AudioPlaybackService{
		fileMap: map[interfaces.AudioEffect]string{
			interfaces.MetronomeClick: "metronome_beat.wav",
			interfaces.SwitchOn:       "switch_on.wav",
			interfaces.SwitchOff:      "switch_off.wav",
			interfaces.Delete:         "delete.wav",
			interfaces.Success:        "success.wav",
		},
	}
}

func (a *AudioPlaybackService) Init() error {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	if err != nil {
		return err
	}
	a.ctx = ctx
	return nil
}

func (a *AudioPlaybackService) Play(effect interfaces.AudioEffect) {
	filename := a.fileMap[effect]
	go a.playFile(filename)
}

func (a *AudioPlaybackService) playFile(filename string) {
	fileBytes, err := resources.AudioFiles.ReadFile("audio_files/" + filename)
	if err != nil {
		logrus.Errorf("Failed to read audio file: %s", err)
		return
	}

	fileBytesReader := bytes.NewReader(fileBytes)
	decoder := wav.NewDecoder(fileBytesReader)
	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		logrus.Errorf("Failed to decode audio: %s", err)
		return
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = uint32(decoder.NumChans)
	deviceConfig.SampleRate = uint32(decoder.SampleRate)

	intBuf := buf.AsIntBuffer().Data
	samples := make([]int16, len(intBuf))
	for i, v := range intBuf {
		samples[i] = int16(v)
	}

	sampleIndex := 0
	done := make(chan struct{})
	var finished sync.Once

	onSendFrames := func(pOutputSample, pInputSamples []byte, framecount uint32) {
		samplesNeeded := int(framecount) * int(decoder.NumChans)
		samplesLeft := len(samples) - sampleIndex

		if samplesLeft <= 0 {
			finished.Do(func() { close(done) })
			return
		}

		if samplesNeeded > samplesLeft {
			samplesNeeded = samplesLeft
		}

		for i := 0; i < samplesNeeded; i++ {
			sample := samples[sampleIndex]
			sampleIndex++
			binary.LittleEndian.PutUint16(pOutputSample[i*2:], uint16(sample))
		}
	}

	device, err := malgo.InitDevice(a.ctx.Context, deviceConfig, malgo.DeviceCallbacks{
		Data: onSendFrames,
	})
	if err != nil {
		logrus.Errorf("Failed to init device: %s", err)
		return
	}

	if err := device.Start(); err != nil {
		logrus.Errorf("Failed to start device: %s", err)
		device.Uninit()
		return
	}

	<-done
	device.Stop()
	device.Uninit()
}

func (a *AudioPlaybackService) Close() error {
	if a.ctx != nil {
		a.ctx.Uninit()
	}
	return nil
}
