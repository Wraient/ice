package internal

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

const upstreamRofiBase = "https://raw.githubusercontent.com/Wraient/ice/main/rofi"

// ensureRofiThemes downloads required .rasi theme files into the user's storage path if absent.
func ensureRofiThemes() {
	cfg := GetGlobalConfig()
	base := os.ExpandEnv(cfg.StoragePath)
	if base == "" {
		return
	}
	_ = os.MkdirAll(base, 0755)
	themes := []string{"selectanime.rasi", "selectanimepreview.rasi", "userinput.rasi"}
	for _, t := range themes {
		dest := filepath.Join(base, t)
		if _, err := os.Stat(dest); err == nil {
			continue
		}
		// download
		resp, err := http.Get(upstreamRofiBase + "/" + t)
		if err != nil {
			continue
		}
		data, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			continue
		}
		_ = os.WriteFile(dest, data, 0644)
	}
}

// cacheImage downloads an image URL and stores it under storagePath/images with hashed name; returns local path.
func cacheImage(imgURL string) string {
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

// BuildRofiOptionLabel optionally prefixes with icon markup if image cached.
// BuildRofiOptionLabel returns a line optionally embedding icon meta for rofi.
// Rofi supports per-entry metadata using NUL separators: 'Text\x00icon\x1f/path/to/icon'.
// We leverage this so the visible text stays 'title' while rofi loads the cached image as icon.
func BuildRofiOptionLabel(title, imageURL string) string {
	local := cacheImage(imageURL)
	if local == "" {
		return title
	}
	return title + "\x00icon\x1f" + local
}

func RofiSelect(options map[string]string, addanimeopt bool) (SelectionOption, error) {
	ensureRofiThemes()
	userCurdConfig := GetGlobalConfig()
	if userCurdConfig.StoragePath == "" {
		userCurdConfig.StoragePath = os.ExpandEnv("${HOME}/.local/share/octo")
	}

	// Create a slice to store the options in the order we want
	var optionsList []string
	for _, value := range options {
		optionsList = append(optionsList, value)
	}

	// Sort the options alphabetically
	sort.Strings(optionsList)

	// Add "Add new anime" and "Quit" options
	if addanimeopt {
		optionsList = append(optionsList, "Add new show", "Quit")
	} else {
		optionsList = append(optionsList, "Quit")
	}

	// Join all options into a single string, separated by newlines
	optionsString := strings.Join(optionsList, "\n")

	// Decide theme: only use preview if previously requested via state flag and we have at least one image marker.
	usePreview := rofiConsumePreviewFlag()
	containsIconMeta := false
	for _, v := range optionsList {
		if strings.Contains(v, "\x00icon\x1f") {
			containsIconMeta = true
			break
		}
	}
	if containsIconMeta {
		usePreview = true
	}
	themeFile := "selectanime.rasi"
	if usePreview {
		themeFile = "selectanimepreview.rasi"
	}
	args := []string{"-dmenu", "-theme", filepath.Join(os.ExpandEnv(userCurdConfig.StoragePath), themeFile), "-i", "-p", "Select a show"}
	if containsIconMeta {
		args = append(args, "-show-icons")
	}
	cmd := exec.Command("rofi", args...)

	// Set up pipes for input and output
	cmd.Stdin = strings.NewReader(optionsString)
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		return SelectionOption{}, fmt.Errorf("failed to run Rofi: %v", err)
	}

	// Get the selected option
	selected := strings.TrimSpace(out.String())

	// Handle special cases
	switch selected {
	case "":
		return SelectionOption{}, fmt.Errorf("no selection made")
	case "Add new show":
		return SelectionOption{Label: "Add new show", Key: "add_new"}, nil
	case "Quit":
		return SelectionOption{Label: "Quit", Key: "-1"}, nil
	}

	// Find the key for the selected value; account for icon metadata (NUL-separated)
	for key, value := range options {
		if value == selected { // exact match
			return SelectionOption{Label: selected, Key: key}, nil
		}
		if idx := strings.IndexByte(value, '\x00'); idx != -1 {
			visible := value[:idx]
			if visible == selected {
				return SelectionOption{Label: selected, Key: key}, nil
			}
		}
	}

	// Fallback: return not found as cancel
	return SelectionOption{Label: "Quit", Key: "-1"}, nil
}

// GetUserInputFromRofi prompts the user for input using Rofi with a custom message
func GetUserInputFromRofi(message string) (string, error) {
	ensureRofiThemes()
	userCurdConfig := GetGlobalConfig()
	if userCurdConfig.StoragePath == "" {
		userCurdConfig.StoragePath = os.ExpandEnv("${HOME}/.local/share/curd")
	}
	// Create the Rofi command
	cmd := exec.Command("rofi", "-dmenu", "-theme", filepath.Join(os.ExpandEnv(userCurdConfig.StoragePath), "userinput.rasi"), "-p", "Input", "-mesg", message)

	// Set up pipes for output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run Rofi: %w", err)
	}

	// Get the entered input
	userInput := strings.TrimSpace(out.String())

	return userInput, nil
}
