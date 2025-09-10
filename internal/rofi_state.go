package internal

import "sync"

var (
	rofiNextUsePreview bool
	rofiStateMu        sync.Mutex
)

// RofiSetNextUsePreview flags the next RofiSelect call to attempt preview (images) theme.
func RofiSetNextUsePreview(v bool) {
	rofiStateMu.Lock()
	rofiNextUsePreview = v
	rofiStateMu.Unlock()
}

func rofiConsumePreviewFlag() bool {
	rofiStateMu.Lock()
	v := rofiNextUsePreview
	rofiNextUsePreview = false
	rofiStateMu.Unlock()
	return v
}
