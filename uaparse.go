package main

import (
	"regexp"
	"strings"
)

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

// deviceType classifies a User-Agent as Mobile, Tablet, or Desktop using
// the same conventions browsers themselves rely on (a bare "Android"
// token without "Mobile" signals a tablet).
func deviceType(ua string) string {
	switch {
	case strings.Contains(ua, "iPad"), strings.Contains(ua, "Android") && !strings.Contains(ua, "Mobile"):
		return "Tablet"
	case strings.Contains(ua, "Mobile"), strings.Contains(ua, "iPhone"), strings.Contains(ua, "Android"):
		return "Mobile"
	default:
		return "Desktop"
	}
}
