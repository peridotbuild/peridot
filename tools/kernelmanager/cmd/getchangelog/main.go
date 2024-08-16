package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	repack_v1 "go.resf.org/peridot/tools/kernelmanager/packager/v1"
)

func readChangelogFromSpec(specBytes []byte) []*repack_v1.ChangelogEntry {
	// Parse spec down to %changelog
	var changelogLines []string

	lines := strings.Split(string(specBytes), "\n")
	for i, line := range lines {
		if strings.Contains(line, "%changelog") {
			changelogLines = lines[i+1:]
			break
		}
	}

	// Parse changelog
	var entries []*repack_v1.ChangelogEntry
	currentEntry := &repack_v1.ChangelogEntry{}
	for _, line := range changelogLines {
		line = strings.TrimSpace(line)
		if line == "" {
			if currentEntry.Subject != "" {
				entries = append(entries, currentEntry)
				currentEntry = &repack_v1.ChangelogEntry{}
			}
			continue
		}

		if strings.HasPrefix(line, "* ") {
			currentEntry.Subject = strings.TrimPrefix(line, "* ")
		} else if strings.HasPrefix(line, "- ") {
			currentEntry.Messages = append(currentEntry.Messages, strings.TrimPrefix(line, "- "))
		}
	}

	if currentEntry.Subject != "" {
		entries = append(entries, currentEntry)
	}

	return entries
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("usage: getchangelog <path to spec>")
	}

	specPath := args[0]
	specBytes, err := os.ReadFile(specPath)
	if err != nil {
		panic(errors.Wrap(err, "failed to read spec file"))
	}

	entries := readChangelogFromSpec(specBytes)

	jsonBytes, err := json.Marshal(entries)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal entries"))
	}

	_, err = fmt.Println(string(jsonBytes))
	if err != nil {
		panic(errors.Wrap(err, "failed to write entries"))
	}
}
