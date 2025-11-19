package models

// ModelExporter is a dummy struct without any functionality. It is used
// to export go structs that would otherwise not get exported by wails
type ModelExporter struct{}

// NoOp is a dummy method that does nothing.
// Binding this method to the app interface makes wails generate JS versions for
// all structs in the method signature
func (m ModelExporter) NoOp(
	_ RecordingRunningEventData,
	_ LiveAudioEventData,
	_ MetronomeStateEventData,
	_ MetronomeBeatEventData,
) {
}
