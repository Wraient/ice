package internal

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	acerBase               = "https://acermovies.val.run"
	endpointSearch         = "/api/search"
	endpointSourceQuality  = "/api/sourceQuality"
	endpointSourceEpisodes = "/api/sourceEpisodes"
	endpointSourceUrl      = "/api/sourceUrl"
)

type AcerSearchResult struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Image string `json:"image"`
}
type AcerQuality struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	EpisodesURL string `json:"episodesUrl"`
	Quality     string `json:"quality"`
}
type AcerEpisode struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func acerPostJSON(path string, payload interface{}, out interface{}) error {
	ensureAcerSession()
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", acerBase+path, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// mimic browser; some endpoints may return empty without these
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:142.0) Gecko/20100101 Firefox/142.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Origin", acerBase)
	req.Header.Set("Referer", acerBase+"/")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-GPC", "1")
	start := time.Now()
	resp, err := getAcerClient().Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if DebugEnabled() {
		DebugLog("ACER REQ path=%s status=%d dur=%s payload=%s cookies=%v", path, resp.StatusCode, time.Since(start), string(body), getAcerClient().Jar.Cookies(req.URL))
		DebugLog("ACER RESP path=%s body=%s", path, truncateForLog(string(raw), 2000))
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("acer %s %d %s", path, resp.StatusCode, string(raw))
	}
	if out == nil {
		return nil
	}
	return json.Unmarshal(raw, out)
}

// truncateForLog limits log size
func truncateForLog(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "...<truncated>"
}

func AcerSearch(query string) ([]AcerSearchResult, error) {
	if DebugEnabled() {
		DebugLog("SEARCH start query=%s", query)
	}
	// add the actual observed key first (searchQuery) then fallbacks
	attempts := []map[string]string{
		{"searchQuery": query},
		{"search": query},
		{"query": query},
		{"title": query},
		{"q": query},
	}
	var lastErr error
	for _, payload := range attempts {
		var resp struct {
			SearchResult []AcerSearchResult `json:"searchResult"`
		}
		if err := acerPostJSON(endpointSearch, payload, &resp); err != nil {
			lastErr = err
			continue
		}
		if len(resp.SearchResult) > 0 {
			if DebugEnabled() {
				DebugLog("SEARCH success payloadKeys=%v results=%d", keysOf(payload), len(resp.SearchResult))
			}
			return resp.SearchResult, nil
		}
	}
	if DebugEnabled() {
		DebugLog("SEARCH no_results query=%s lastErr=%v", query, lastErr)
	}
	if lastErr != nil {
		return nil, lastErr
	}
	return []AcerSearchResult{}, nil
}

func keysOf(m map[string]string) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
func AcerGetQualities(url string) ([]AcerQuality, error) {
	var resp struct {
		SourceQualityList []AcerQuality `json:"sourceQualityList"`
	}
	if err := acerPostJSON(endpointSourceQuality, map[string]string{"url": url}, &resp); err != nil {
		return nil, err
	}
	return resp.SourceQualityList, nil
}
func AcerGetEpisodes(url string) ([]AcerEpisode, error) {
	var resp struct {
		SourceEpisodes []AcerEpisode `json:"sourceEpisodes"`
	}
	if err := acerPostJSON(endpointSourceEpisodes, map[string]string{"url": url}, &resp); err != nil {
		return nil, err
	}
	return resp.SourceEpisodes, nil
}
func AcerGetSource(intermediate, seriesType string) (string, error) {
	var resp struct {
		SourceUrl string `json:"sourceUrl"`
	}
	if err := acerPostJSON(endpointSourceUrl, map[string]string{"url": intermediate, "seriesType": seriesType}, &resp); err != nil {
		return "", err
	}
	if resp.SourceUrl == "" {
		return "", errors.New("empty source url")
	}
	return resp.SourceUrl, nil
}

func AcerSelectQuality(qualities []AcerQuality) (AcerQuality, bool) {
	cfg := GetGlobalConfig()
	pref := strings.ToLower(cfg.PreferredQuality)
	if pref != "" && len(qualities) > 1 { // only auto-pick if single unambiguous
		var exact []AcerQuality
		for _, q := range qualities {
			if strings.EqualFold(q.Quality, pref) {
				exact = append(exact, q)
			}
		}
		if len(exact) == 1 {
			return exact[0], false
		}
		if len(exact) > 1 {
			qualities = exact
		} // narrow but still need selection
	}
	if len(qualities) == 1 {
		return qualities[0], false
	}
	return AcerQuality{}, true
}
func AcerQualityOptions(qs []AcerQuality) map[string]string {
	m := map[string]string{}
	for i, q := range qs {
		season := extractSeason(q.Title)
		if season == 0 {
			season = 1
		}
		m[fmt.Sprintf("%d", i)] = fmt.Sprintf("S%02d | %s | %s", season, q.Quality, q.Title)
	}
	return m
}

var epNumRe = regexp.MustCompile(`(?i)episode\s+(\d+)`)
var seasonRe = regexp.MustCompile(`(?i)season\s*(\d+)`)

// --- session / cookie management ---
var acerClient *http.Client
var acerOnce sync.Once

func getAcerClient() *http.Client {
	acerOnce.Do(func() {
		jar, _ := cookiejar.New(nil)
		acerClient = &http.Client{Timeout: 30 * time.Second, Jar: jar}
	})
	return acerClient
}

func ensureAcerSession() {
	// If we already have a visitor_id cookie, skip
	u, _ := url.Parse(acerBase)
	if u == nil {
		return
	}
	cookies := getAcerClient().Jar.Cookies(u)
	for _, c := range cookies {
		if c.Name == "visitor_id" && c.Value != "" {
			return
		}
	}
	// Fetch root to obtain visitor cookie
	req, _ := http.NewRequest("GET", acerBase+"/", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:110.0) Gecko/20100101 Firefox/110.0")
	resp, err := getAcerClient().Do(req)
	if err != nil {
		if DebugEnabled() {
			DebugLog("ACER session init error=%v", err)
		}
		return
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	// After fetch, check again
	cookies = getAcerClient().Jar.Cookies(u)
	hasVisitor := false
	for _, c := range cookies {
		if c.Name == "visitor_id" && c.Value != "" {
			hasVisitor = true
			break
		}
	}
	if !hasVisitor {
		// generate our own visitor_id (uuid-like) 8-4-4-4-12 hex
		buf := make([]byte, 16)
		if _, err := rand.Read(buf); err == nil {
			id := hex.EncodeToString(buf)
			// format
			visitor := fmt.Sprintf("%s-%s-%s-%s-%s", id[0:8], id[8:12], id[12:16], id[16:20], id[20:32])
			getAcerClient().Jar.SetCookies(u, []*http.Cookie{{
				Name:  "visitor_id",
				Value: visitor,
				Path:  "/",
				// set a long expiry
				Expires: time.Now().Add(30 * 24 * time.Hour),
			}})
			if DebugEnabled() {
				DebugLog("ACER session synthesized visitor_id=%s", visitor)
			}
		}
	}
	if DebugEnabled() {
		DebugLog("ACER session cookies=%v", getAcerClient().Jar.Cookies(u))
	}
}

func atoiSafe(s string) int { var n int; fmt.Sscanf(s, "%d", &n); return n }
func parseEpNum(t string) int {
	if m := epNumRe.FindStringSubmatch(t); len(m) == 2 {
		return atoiSafe(m[1])
	}
	return 0
}
func extractSeason(t string) int {
	if m := seasonRe.FindStringSubmatch(t); len(m) == 2 {
		return atoiSafe(m[1])
	}
	return 0
}

func AcerBuildShow(showURL string, q AcerQuality) (*Show, error) {
	if q.EpisodesURL == "" {
		return &Show{Id: showURL, Name: showURL, EpisodesList: []EpisodeEntry{{Name: q.Title, ID: q.URL, Episode: 1, Season: 1}}}, nil
	}
	eps, err := AcerGetEpisodes(q.EpisodesURL)
	if err != nil {
		return nil, err
	}
	season := extractSeason(q.Title)
	if season == 0 {
		season = 1
	}
	list := make([]EpisodeEntry, 0, len(eps))
	for _, e := range eps {
		list = append(list, EpisodeEntry{Name: e.Title, ID: e.Link, Episode: parseEpNum(e.Title), Season: season})
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Episode < list[j].Episode })
	return &Show{Id: showURL, Name: showURL, EpisodesList: list}, nil
}

// GetShow now delegates to GetShowWithQualitySeason with no forced season/quality.
func GetShow(id string) (*Show, error) { return GetShowWithQualitySeason(id, 0, "") }

// GetShowWithQualitySeason allows specifying season (if known) and a quality preference to avoid re-prompting.
func GetShowWithQualitySeason(id string, season int, qualityPref string) (*Show, error) {
	qualities, err := AcerGetQualities(id)
	if err != nil {
		return nil, err
	}
	if len(qualities) == 0 {
		return nil, fmt.Errorf("no qualities returned")
	}
	seasonMap := map[int][]AcerQuality{}
	var seasons []int
	for _, q := range qualities {
		s := extractSeason(q.Title)
		if s == 0 {
			s = 1
		}
		if _, ok := seasonMap[s]; !ok {
			seasons = append(seasons, s)
		}
		seasonMap[s] = append(seasonMap[s], q)
	}
	sort.Ints(seasons)
	chosenSeason := season
	if chosenSeason == 0 { // need selection if multiple seasons
		if len(seasons) > 1 {
			opts := map[string]string{}
			for _, s := range seasons {
				opts[fmt.Sprintf("%d", s)] = fmt.Sprintf("Season %d", s)
			}
			sel, err := DynamicSelect(opts)
			if err != nil || sel.Key == "-1" {
				return nil, fmt.Errorf("season selection aborted")
			}
			chosenSeason = atoiSafe(sel.Key)
		} else {
			chosenSeason = seasons[0]
		}
	}
	list := seasonMap[chosenSeason]
	if qualityPref != "" { // filter list to requested quality if exists
		var filtered []AcerQuality
		for _, q := range list {
			if strings.EqualFold(q.Quality, qualityPref) {
				filtered = append(filtered, q)
			}
		}
		if len(filtered) > 0 {
			list = filtered
		}
	}
	selQ, need := AcerSelectQuality(list)
	if need {
		opts := AcerQualityOptions(list)
		chosen, err := DynamicSelect(opts)
		if err != nil || chosen.Key == "-1" {
			return nil, fmt.Errorf("quality selection aborted")
		}
		idx := atoiSafe(chosen.Key)
		if idx < 0 || idx >= len(list) {
			return nil, fmt.Errorf("invalid quality index")
		}
		selQ = list[idx]
		if qualityPref == "" {
			cfg := GetGlobalConfig()
			cfg.PreferredQuality = selQ.Quality
			_ = SaveGlobalConfig()
		}
	}
	show, err := AcerBuildShow(id, selQ)
	if show != nil {
		for i := range show.EpisodesList {
			if show.EpisodesList[i].Season == 0 {
				show.EpisodesList[i].Season = extractSeason(selQ.Title)
				if show.EpisodesList[i].Season == 0 {
					show.EpisodesList[i].Season = chosenSeason
				}
			}
		}
	}
	return show, err
}
