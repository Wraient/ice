# Ice

Ice is the evolution of the earlier "Octopus" terminal player. It lets you search, select, and stream shows or movies directly from the terminal (or via Rofi) using `mpv`, while tracking progress and resuming seamlessly. The content provider has moved from the old vadapav API to the Acer Movies workflow.

## Highlights
* Acer Movies integration (multi-key search fallback + session cookie handling)
* Season & quality selection (with remembered preferred quality)
* Episode picker & automatic next‑episode advance after completion threshold
* Resume playback (per show stored in a simple CSV database)
* MPV integration with IPC (time-pos, duration, resume seek)
* Clean media titles passed to mpv (`--force-media-title`)
* Rofi or pure terminal interactive selection
* Debug logging to `~/.config/ice/debug.log` (HTTP payloads & responses truncated for safety)

## Current Status
Focused on Acer-only provider. Module path rename is pending; code may still import under the old module name internally, but the binary is intended to be used as `ice`.

## Install
You need: `mpv` (required), `rofi` (optional), a working Go toolchain if building from source.

### Build from Source
```bash
git clone https://github.com/Wraient/ice.git
cd ice
go build -o ice ./cmd/ice
./ice --help
```

### Simple Global Install (Go >=1.21)
```bash
go install github.com/Wraient/ice/cmd/ice@latest
```

Ensure `GOBIN` (typically `$HOME/go/bin`) is on your PATH.

### Dependencies
Arch (pacman): `sudo pacman -S mpv rofi`
Debian/Ubuntu: `sudo apt install mpv rofi`
Fedora: `sudo dnf install mpv rofi`
openSUSE: `sudo zypper install mpv rofi`

Rofi is optional; omit if you only want terminal selection.

## Usage
Basic run (interactive search prompt in terminal):
```bash
ice
```

Rofi mode:
```bash
ice -rofi
```

Resume or search selection appears if you already have stored shows.

### Flags
| Flag | Description | Default |
|------|-------------|---------|
| `-player` | Media player (mpv only for now) | `mpv` |
| `-storage-path` | Base data dir (db + logs) | `$HOME/.local/share/ice` |
| `-percentage-to-mark-complete` | Percent watched before auto-advancing | `92` |
| `-rofi` | Force rofi selection menus | `false` |
| `-no-rofi` | Force disable rofi (even if config enables) | `false` |
| `-save-mpv-speed` | Persist last mpv playback speed | `true` |
| `-next-episode-prompt` | (Reserved / legacy) ask before advancing | `false` |
| `-quality` | Override preferred quality for this run | (empty) |
| `-debug` | Write verbose provider + request logs | `false` |

### Flow
1. Start `ice` → choose Continue or Search.
2. Enter query (terminal) or rofi prompt.
3. Search uses a fallback sequence of JSON keys: `searchQuery`, `search`, `query`, `title`, `q` until results appear.
4. Select a result → fetch qualities grouped by season.
5. If multiple seasons: prompted to pick one.
6. If multiple qualities: auto-picks preferred or prompts.
7. If multiple episodes: choose one (or when resuming you can change episode).
8. mpv launches; progress is updated every second.
9. On completion threshold, next episode auto-plays (if any).

### Resume Data
Stored in CSV: `shows.db` at `$STORAGE_PATH` with columns:
```
ShowID,EpisodeID,PlaybackTime,Season,Quality,EpisodeNum
```
ShowID = original Acer content page URL (not the intermediate source URL). This lets the app re-fetch updated episode lists.

### Debugging
Enable with:
```bash
ice --debug
```
Writes to `~/.config/ice/debug.log` including:
* Session init & synthesized visitor cookie
* Search request/response bodies (truncated)
* Endpoint timing & payloads

### Media Titles
Episode names are normalized: slug / quality junk words are stripped; mpv window & OSD show `Show Name - SxxExx Episode Title`.

## Configuration File
Created at: `~/.config/ice/ice.conf`
Keys (auto-added if missing):
```
Player=mpv
StoragePath=$HOME/.local/share/ice
PercentageToMarkComplete=92
NextEpisodePrompt=false
RofiSelection=false
SaveMpvSpeed=true
PreferredQuality=1080p
```
Editing can be done manually; the app updates `PreferredQuality` when you pick a new one.

## Architecture Overview
| Component | Purpose |
|-----------|---------|
| `internal/acer.go` | Acer API client (search, qualities, episodes, source URL) + session & cookie logic |
| `internal/player.go` | mpv spawn & IPC (seek, duration, speed, title) |
| `internal/database.go` | CSV persistence for progress/resume |
| `internal/naming.go` | Slug → pretty title normalization |
| `internal/debug.go` | Debug file logging facility |
| `cmd/ice/main.go` | CLI orchestration (flags, selection flow) |

## Roadmap / TODO
* Refine per-show remembered quality instead of global.
* Module path rename cleanup.
* Richer episode metadata extraction.
* Optional parallel prefetch of next episode URL.

## License
MIT – see [LICENSE](LICENSE).

## Credits
* Original Octopus concept foundation.
* mpv project for robust IPC.

---
Feedback & issues welcome via repository tracker.
