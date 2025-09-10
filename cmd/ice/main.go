package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/Wraient/octopus/internal"
)

func main() {
	var user internal.User
	var show internal.TVShow

	var homeDir string
	if runtime.GOOS == "windows" {
		homeDir = os.Getenv("USERPROFILE")
	} else {
		homeDir = os.Getenv("HOME")
	}
	configFilePath := filepath.Join(homeDir, ".config", "ice", "ice.conf")

	userIceConfig, err := internal.LoadConfig(configFilePath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	internal.SetGlobalConfig(&userIceConfig)

	databaseFile := filepath.Join(os.ExpandEnv(userIceConfig.StoragePath), "shows.db")
	logFile := filepath.Join(os.ExpandEnv(userIceConfig.StoragePath), "debug.log")

	// Flags
	flag.StringVar(&userIceConfig.Player, "player", userIceConfig.Player, "Player to use (mpv)")
	flag.StringVar(&userIceConfig.StoragePath, "storage-path", userIceConfig.StoragePath, "Path to storage")
	flag.IntVar(&userIceConfig.PercentageToMarkComplete, "percentage-to-mark-complete", userIceConfig.PercentageToMarkComplete, "Percentage to mark complete")
	flag.BoolVar(&userIceConfig.SaveMpvSpeed, "save-mpv-speed", userIceConfig.SaveMpvSpeed, "Save MPV speed")
	flag.BoolVar(&userIceConfig.NextEpisodePrompt, "next-episode-prompt", userIceConfig.NextEpisodePrompt, "Prompt next episode")
	preferredQuality := flag.String("quality", "", "Override preferred quality for this run")
	rofiSelection := flag.Bool("rofi", false, "Use rofi for selection")
	noRofi := flag.Bool("no-rofi", false, "Disable rofi selection")
	debugFlag := flag.Bool("debug", false, "Enable verbose debug logging to debug.log in config directory")
	flag.Parse()

	if *preferredQuality != "" {
		userIceConfig.PreferredQuality = *preferredQuality
	}
	if *rofiSelection {
		userIceConfig.RofiSelection = true
	}
	if *noRofi || runtime.GOOS == "windows" {
		userIceConfig.RofiSelection = false
	}

	// Save config if modified via flags
	internal.SetGlobalConfig(&userIceConfig)
	if *debugFlag {
		configDir := filepath.Dir(configFilePath)
		debugFile := filepath.Join(configDir, "debug.log")
		if err := internal.EnableDebug(debugFile); err == nil {
			internal.OctoOut("Debug file: " + debugFile)
		} else {
			internal.OctoOut("Failed to enable debug: " + err.Error())
		}
	}
	_ = internal.SaveGlobalConfig()

	internal.ClearLog(logFile)

	// Continue watching logic
	shows := internal.LocalGetAllShows(databaseFile)
	if len(shows) > 0 {
		options := map[string]string{"n": "Search new", "y": "Continue watching"}
		sel, err := internal.DynamicSelect(options)
		if err == nil && sel.Key == "y" {
			showOptions := map[string]string{}
			for _, s := range shows {
				friendly := internal.PrettyShowName(s.ID)
				label := friendly
				if s.EpisodeNum > 0 && s.Season > 0 {
					label = fmt.Sprintf("%s S%02dE%02d", friendly, s.Season, s.EpisodeNum)
				}
				if s.PlaybackTime > 0 {
					label += " (resume)"
				}
				showOptions[s.ID] = label
			}
			picked, err := internal.DynamicSelect(showOptions)
			if err == nil && picked.Key != "-1" {
				for _, s := range shows {
					if s.ID == picked.Key {
						show = s
						user.Resume = s.PlaybackTime > 0
						break
					}
				}
			}
		}
	}
	// If we chose an existing show we may want to allow picking another episode before starting

	// Fresh search if needed
	var query string
	if show.ID == "" {
		if userIceConfig.RofiSelection {
			q, err := internal.GetUserInputFromRofi("Enter name:")
			if err != nil {
				internal.ExitOcto("", err)
				return
			}
			query = q
		} else {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter name: ")
			q, _ := reader.ReadString('\n')
			query = strings.TrimSpace(q)
		}
		results, err := internal.AcerSearch(query)
		if err != nil {
			internal.ExitOcto("", err)
			return
		}
		if len(results) == 0 {
			internal.ExitOcto(fmt.Sprintf("No results: %s", query), nil)
		}
		opt := map[string]string{}
		for _, r := range results {
			opt[r.URL] = r.Title
		}
		picked, err := internal.DynamicSelect(opt)
		if err != nil || picked.Key == "-1" {
			return
		}
		show = internal.TVShow{ID: picked.Key, EpisodeID: "", PlaybackTime: 0}
	}

	// Determine movie vs series & pick quality/episode
	// Use GetShowWithQualitySeason to respect saved season/quality when resuming
	forcedSeason := 0
	qualityPref := ""
	if show.Season > 0 {
		forcedSeason = show.Season
	}
	if show.Quality != "" {
		qualityPref = show.Quality
	}
	acerShow, err := internal.GetShowWithQualitySeason(show.ID, forcedSeason, qualityPref)
	if err != nil {
		internal.ExitOcto("", err)
		return
	}

	// If resuming, keep existing episode ID
	if !user.Resume || show.EpisodeID == "" {
		// If movie (single episode) or user not resuming, select episode if multiple
		if len(acerShow.EpisodesList) > 1 && !user.Resume {
			epOpts := map[string]string{}
			for _, ep := range acerShow.EpisodesList {
				epOpts[ep.ID] = fmt.Sprintf("S%02dE%02d %s", ep.Season, ep.Episode, ep.Name)
			}
			selectedEp, err := internal.DynamicSelect(epOpts)
			if err != nil || selectedEp.Key == "-1" {
				return
			}
			show.EpisodeID = selectedEp.Key
		} else {
			show.EpisodeID = acerShow.EpisodesList[0].ID
		}
	} else if user.Resume && len(acerShow.EpisodesList) > 1 { // Offer to pick a different episode
		pickOpts := map[string]string{"r": "Resume current episode"}
		for _, ep := range acerShow.EpisodesList {
			marker := ""
			if ep.ID == show.EpisodeID { marker = " *" }
			pickOpts[ep.ID] = fmt.Sprintf("S%02dE%02d %s%s", ep.Season, ep.Episode, ep.Name, marker)
		}
		choice, err := internal.DynamicSelect(pickOpts)
		if err == nil && choice.Key != "-1" && choice.Key != "r" {
			show.EpisodeID = choice.Key
			show.PlaybackTime = 0
			user.Resume = false
		}
	}
	// Persist season/quality metadata from built show (episodes share season; we infer from first)
	if len(acerShow.EpisodesList) > 0 {
		show.Season = acerShow.EpisodesList[0].Season
	}
	show.Quality = internal.GetGlobalConfig().PreferredQuality
	// Episode number for progression
	for _, ep := range acerShow.EpisodesList {
		if ep.ID == show.EpisodeID {
			show.EpisodeNum = ep.Episode
			break
		}
	}

	// Acquire playback URL for current episode/movie
	seriesType := "episode"
	if len(acerShow.EpisodesList) == 1 {
		seriesType = "movie"
	}
	playbackURL, err := internal.AcerGetSource(show.EpisodeID, seriesType)
	if err != nil {
		internal.ExitOcto("", err)
		return
	}

	// Build media title for mpv: derive show name + episode label
	mediaTitle := internal.PrettyShowName(show.ID)
	epLabel := ""
	for _, ep := range acerShow.EpisodesList {
		if ep.ID == show.EpisodeID {
			epLabel = fmt.Sprintf("S%02dE%02d %s", ep.Season, ep.Episode, ep.Name)
			break
		}
	}
	if epLabel != "" {
		mediaTitle = fmt.Sprintf("%s - %s", mediaTitle, epLabel)
	}
	internal.OctoOut(fmt.Sprintf("Playing %s", mediaTitle))
	user.Player.SocketPath, err = internal.PlayWithMPV(playbackURL, mediaTitle)
	if err != nil {
		internal.ExitOcto("", err)
		return
	}

	for {
		// Duration goroutine
		go func() {
			for {
				if user.Player.Started && user.Player.Duration == 0 {
					if d, err := internal.MPVSendCommand(user.Player.SocketPath, []interface{}{"get_property", "duration"}); err == nil {
						if dv, ok := d.(float64); ok {
							user.Player.Duration = int(dv + 0.5)
						}
					}
					break
				}
				time.Sleep(time.Second)
			}
		}()

		// Playback monitor loop
	skipLoop:
		for {
			time.Sleep(time.Second)
			pos, err := internal.MPVSendCommand(user.Player.SocketPath, []interface{}{"get_property", "time-pos"})
			if err != nil {
				if user.Player.Started {
					percentage := internal.PercentageWatched(show.PlaybackTime, user.Player.Duration)
					if percentage >= float64(userIceConfig.PercentageToMarkComplete) { // get next episode
						nshow, err := internal.GetShowWithQualitySeason(show.ID, show.Season, show.Quality)
						if err == nil {
							nextEp := internal.GetNextEpisode(nshow, show.EpisodeID)
							if nextEp != nil {
								show.EpisodeID = nextEp.ID
								show.PlaybackTime = 0
								show.EpisodeNum = nextEp.Episode
								break skipLoop
							}
						}
						internal.ExitOcto("", nil)
					} else {
						internal.ExitOcto("", nil)
					}
				}
				break skipLoop
			}
			if pos != nil {
				if !user.Player.Started {
					user.Player.Started = true
				}
				if user.Resume {
					internal.SeekMPV(user.Player.SocketPath, show.PlaybackTime)
					user.Resume = false
				}
				if fv, ok := pos.(float64); ok {
					show.PlaybackTime = int(fv + 0.5)
					_ = internal.LocalUpdateShow(databaseFile, show)
				}
			}
		}
		if show.PlaybackTime == 0 { // next episode
			user.Player.Duration = 0
			user.Player.Started = false
			// fetch fresh playback url
			playbackURL, err = internal.AcerGetSource(show.EpisodeID, "episode")
			if err != nil {
				internal.ExitOcto("", err)
			}
			// Update media title for next episode
			mediaTitle = internal.PrettyShowName(show.ID)
			epLabel = ""
			for _, ep := range acerShow.EpisodesList {
				if ep.ID == show.EpisodeID {
					epLabel = fmt.Sprintf("S%02dE%02d %s", ep.Season, ep.Episode, ep.Name)
					break
				}
			}
			if epLabel != "" {
				mediaTitle = fmt.Sprintf("%s - %s", mediaTitle, epLabel)
			}
			user.Player.SocketPath, err = internal.PlayWithMPV(playbackURL, mediaTitle)
			if err != nil {
				internal.ExitOcto("", err)
			}
		}
	}
}
