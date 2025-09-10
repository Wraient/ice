package internal

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// CacheImage downloads and stores an image locally returning its local path or empty string on failure.
func CacheImage(imgURL string) string {
	if imgURL == "" {
		return ""
	}
	cfg := GetGlobalConfig()
	base := os.ExpandEnv(cfg.StoragePath)
	if base == "" {
		return ""
	}
	imgDir := filepath.Join(base, "images")
	_ = os.MkdirAll(imgDir, 0755)
	h := sha1.Sum([]byte(imgURL))
	name := fmt.Sprintf("%x.jpg", h[:8])
	dest := filepath.Join(imgDir, name)
	if _, err := os.Stat(dest); err == nil {
		return dest
	}
	resp, err := http.Get(imgURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return ""
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	_ = os.WriteFile(dest, data, 0644)
	return dest
}
