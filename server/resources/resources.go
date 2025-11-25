package resources

import "embed"

//go:embed audio_files/*.wav
var AudioFiles embed.FS
