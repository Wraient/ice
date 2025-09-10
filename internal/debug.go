package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var debugMode bool
var debugFilePath string
var debugMu sync.Mutex

// EnableDebug sets debug mode and initializes / truncates the debug log file.
func EnableDebug(logPath string) error {
	debugMu.Lock()
	defer debugMu.Unlock()
	debugMode = true
	debugFilePath = os.ExpandEnv(logPath)
	if err := os.MkdirAll(filepath.Dir(debugFilePath), 0755); err != nil {
		return err
	}
	f, err := os.Create(debugFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, _ = f.WriteString(fmt.Sprintf("[DEBUG] session start %s\n", time.Now().Format(time.RFC3339)))
	return nil
}

func DebugEnabled() bool { return debugMode }

// DebugLog writes a formatted line to debug log if enabled.
func DebugLog(format string, args ...interface{}) {
	if !debugMode || debugFilePath == "" {
		return
	}
	debugMu.Lock()
	defer debugMu.Unlock()
	f, err := os.OpenFile(debugFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	ts := time.Now().Format("2006-01-02 15:04:05")
	line := fmt.Sprintf(format, args...)
	fmt.Fprintf(f, "%s %s\n", ts, line)
}

func GetDebugFilePath() string { return debugFilePath }
