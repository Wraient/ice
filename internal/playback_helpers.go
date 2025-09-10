package internal

// GetNextEpisode returns the next episode entry after currentEpisodeID within the Show list.
func GetNextEpisode(currentShow *Show, currentEpisodeID string) *EpisodeEntry {
	if currentShow == nil {
		return nil
	}
	idx := -1
	for i, ep := range currentShow.EpisodesList {
		if ep.ID == currentEpisodeID {
			idx = i
			break
		}
	}
	if idx >= 0 && idx < len(currentShow.EpisodesList)-1 {
		return &currentShow.EpisodesList[idx+1]
	}
	return nil
}
