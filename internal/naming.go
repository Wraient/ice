package internal

import (
    "net/url"
    "path"
    "regexp"
    "strings"
    "unicode"
)

var qualityTokens = []string{"480p","720p","1080p","2160p","4k","x264","x265","bluray","webrip","web-dl","hdrip"}
var junkTokens = []string{"download","watch","stream","dual","audio","english","hindi","season","complete"}
var multiDashRe = regexp.MustCompile(`[-_]+`)
var spaceMultiRe = regexp.MustCompile(`\s+`)
var episodePatternRe = regexp.MustCompile(`(?i)s(\d{1,2})e(\d{1,3})`)

// PrettyShowName attempts to convert an Acer slug/URL or file-like name into a nicer title.
func PrettyShowName(raw string) string {
    candidate := raw
    if strings.HasPrefix(raw, "http") {
        if u, err := url.Parse(raw); err == nil {
            segs := strings.Split(strings.Trim(u.Path, "/"), "/")
            if len(segs) > 0 && segs[len(segs)-1] != "" { candidate = segs[len(segs)-1] }
        }
    }
    if ext := path.Ext(candidate); len(ext) > 0 && len(ext) <= 5 { candidate = strings.TrimSuffix(candidate, ext) }
    candidate = strings.ReplaceAll(candidate, "%20", " ")
    candidate = multiDashRe.ReplaceAllString(candidate, " ")
    candidate = strings.ReplaceAll(candidate, ".", " ")
    candidate = strings.ReplaceAll(candidate, "_", " ")
    lower := strings.ToLower(candidate)
    for _, t := range append(qualityTokens, junkTokens...) {
        lower = strings.ReplaceAll(lower, t, " ")
    }
    lower = spaceMultiRe.ReplaceAllString(lower, " ")
    lower = strings.TrimSpace(lower)
    if lower == "" { return raw }
    words := strings.Fields(lower)
    for i, w := range words { words[i] = titleCaseWord(w) }
    title := strings.Join(words, " ")
    // Normalize episode tokens to upper SxxEyy if present
    title = episodePatternRe.ReplaceAllStringFunc(title, func(s string) string { return strings.ToUpper(s) })
    return title
}

func titleCaseWord(w string) string {
    small := map[string]bool{"and":true, "or":true, "the":true, "of":true, "a":true, "an":true}
    if small[w] { return w }
    r := []rune(w)
    if len(r) == 0 { return w }
    r[0] = unicode.ToUpper(r[0])
    return string(r)
}
