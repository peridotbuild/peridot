package kernelorg

import (
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

type releaseAtomItem struct {
	Title string `xml:"title"`
}

type releaseAtomChannel struct {
	Items []*releaseAtomItem `xml:"item"`
}

type releaseAtom struct {
	// RSS feed
	// https://www.kernel.org/feeds/kdist.xml
	XMLName xml.Name            `xml:"rss"`
	Channel *releaseAtomChannel `xml:"channel"`
}

const atomURL = "https://www.kernel.org/feeds/kdist.xml"

var (
	ErrNoRelease = errors.New("no release")
	ErrNotFound  = errors.New("not found")
)

func GetLatestMLVersion() (string, error) {
	f, err := http.Get(atomURL)
	if err != nil {
		return "", err
	}
	defer f.Body.Close()

	var atom releaseAtom
	err = xml.NewDecoder(f.Body).Decode(&atom)
	if err != nil {
		return "", err
	}

	if atom.Channel == nil {
		return "", ErrNoRelease
	}

	if len(atom.Channel.Items) == 0 {
		return "", ErrNoRelease
	}

	for _, item := range atom.Channel.Items {
		if strings.HasSuffix(item.Title, ": mainline") {
			return strings.TrimSuffix(item.Title, ": mainline"), nil
		}
	}

	return "", ErrNotFound
}

func GetLatestStableVersion() (string, error) {
	f, err := http.Get(atomURL)
	if err != nil {
		return "", err
	}
	defer f.Body.Close()

	var atom releaseAtom
	err = xml.NewDecoder(f.Body).Decode(&atom)
	if err != nil {
		return "", err
	}

	if atom.Channel == nil {
		return "", ErrNoRelease
	}

	if len(atom.Channel.Items) == 0 {
		return "", ErrNoRelease
	}

	for _, item := range atom.Channel.Items {
		if strings.HasSuffix(item.Title, ": stable") {
			return strings.TrimSuffix(item.Title, ": stable"), nil
		}
	}

	return "", ErrNotFound
}

func GetLTVersion(majorVersion string) (string, error) {
	f, err := http.Get(atomURL)
	if err != nil {
		return "", err
	}
	defer f.Body.Close()

	var atom releaseAtom
	err = xml.NewDecoder(f.Body).Decode(&atom)
	if err != nil {
		return "", err
	}

	if atom.Channel == nil {
		return "", ErrNoRelease
	}

	if len(atom.Channel.Items) == 0 {
		return "", ErrNoRelease
	}

	for _, item := range atom.Channel.Items {
		if strings.HasPrefix(item.Title, majorVersion) && strings.HasSuffix(item.Title, ": longterm") {
			return strings.TrimSuffix(item.Title, ": longterm"), nil
		}
	}

	return "", ErrNotFound
}
