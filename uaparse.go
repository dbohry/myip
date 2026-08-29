package main

import "regexp"

var browserPatterns = []struct {
	name string
	re   *regexp.Regexp
}{
	{"Edge", regexp.MustCompile(`Edg/([\d.]+)`)},
	{"Opera", regexp.MustCompile(`(?:OPR|Opera)/([\d.]+)`)},
	{"Firefox", regexp.MustCompile(`Firefox/([\d.]+)`)},
	{"Chrome", regexp.MustCompile(`Chrome/([\d.]+)`)},
	{"Safari", regexp.MustCompile(`Version/([\d.]+).*Safari`)},
}

var osPatterns = []struct {
	name string
	re   *regexp.Regexp
}{
	{"Windows", regexp.MustCompile(`Windows NT`)},
	{"iOS", regexp.MustCompile(`i(?:Phone|Pad|Pod)`)},
	{"macOS", regexp.MustCompile(`Mac OS X`)},
	{"Android", regexp.MustCompile(`Android`)},
	{"Linux", regexp.MustCompile(`Linux`)},
}

// parseUserAgent extracts a best-effort browser and OS name from a
// User-Agent string. It is a small heuristic parser, not a full UA
// database, since exact versions/devices aren't needed here.
func parseUserAgent(ua string) (browser, os string) {
	browser = "Unknown"
	for _, p := range browserPatterns {
		if m := p.re.FindStringSubmatch(ua); m != nil {
			browser = p.name + " " + m[1]
			break
		}
	}

	os = "Unknown"
	for _, p := range osPatterns {
		if p.re.MatchString(ua) {
			os = p.name
			break
		}
	}

	return browser, os
}
