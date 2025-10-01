package models


type ModelExporter struct{}

// Export is a dummy method that does nothing.
// Binding this method to the app interface makes wails generate JS versions for 
// all structs in the method signature
func (m ModelExporter) NoOp(
	_ RunningEventData,
	_ LiveAudioEventData,
) {
}
