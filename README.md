# Ice

Watch Shows/Movies with Tracking and resume.

## Demo

https://github.com/user-attachments/assets/7c6da538-c780-4ef7-8b57-31fbd3f01c22

## Highlights
- Fast speed since links are from googleusercontent (20 Mb/s+)
- Large collection
- Track Shows/Movies to resume easily

## Limitations
- Cannot seek through Shows/Movies since the link is not seekable.

## Install
You need: `mpv` (required), `rofi` (optional), a working Go toolchain if building from source.

<details>
<summary>Generic Installation</summary>

Choose the appropriate binary for your system:

```bash
# For Linux x86_64:
curl -Lo ice https://github.com/Wraient/ice/releases/latest/download/ice-linux-amd64

# For Linux ARM64:
curl -Lo ice https://github.com/Wraient/ice/releases/latest/download/ice-linux-arm64

# For macOS ARM64:
curl -Lo ice https://github.com/Wraient/ice/releases/latest/download/ice-darwin-arm64

# For macOS x86_64:
curl -Lo ice https://github.com/Wraient/ice/releases/latest/download/ice-darwin-amd64

chmod +x ice
sudo mv ice /usr/bin/
ice
```
</details>

<details>
<summary>Build from Source</summary>
  
```bash
git clone https://github.com/Wraient/ice.git
cd ice
go build -o ice ./cmd/ice
./ice --help
```

</details>

<details>
  
<summary>Simple Global Install (Go >=1.21)</summary>

```bash
go install github.com/Wraient/ice/cmd/ice@latest
```

Ensure `GOBIN` (typically `$HOME/go/bin`) is on your PATH.

</details>

## Dependencies
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
| `-quality` | Override preferred quality for this run |  |
| `-debug` | Write verbose provider + request logs | `false` |

Example

```bash
ice -player=mpv
ice -rofi
ice -debug
ice -quality 1080p
```

### Debugging
Enable with:
```bash
ice -debug
```
Writes to `~/.config/ice/debug.log` including:
* Session init & synthesized visitor cookie
* Search request/response bodies (truncated)
* Endpoint timing & payloads

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

## License
[MIT LICENSE](LICENSE)

## Credits
- https://acermovies.val.run/ (For api)
- [Jerry](https://github.com/justchokingaround/jerry) (For inspiration)
