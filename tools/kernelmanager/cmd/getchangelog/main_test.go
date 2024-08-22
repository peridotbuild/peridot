package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func readSpec(t *testing.T, name string) []byte {
	specBytes, err := os.ReadFile(fmt.Sprintf("testdata/%s.spec", name))
	require.Nil(t, err)

	return specBytes
}

func TestReadChangelogFromSpec_1(t *testing.T) {
	specBytes := readSpec(t, "1")
	entries := readChangelogFromSpec(specBytes)

	require.Len(t, entries, 1)
	require.Equal(t, "Tue Feb 22 2024 Mustafa Gezen - 1.0-123", entries[0].Subject)
	require.Len(t, entries[0].Messages, 1)
	require.Equal(t, "Test changelog", entries[0].Messages[0])
}

func TestReadChangelogFromSpec_2(t *testing.T) {
	specBytes := readSpec(t, "2")
	entries := readChangelogFromSpec(specBytes)

	require.Len(t, entries, 2)

	require.Equal(t, "Tue Feb 22 2024 Mustafa Gezen - 1.0-123", entries[0].Subject)
	require.Len(t, entries[0].Messages, 1)
	require.Equal(t, "Test changelog", entries[0].Messages[0])

	require.Equal(t, "Tue Feb 22 2024 Mustafa Gezen - 1.0-123", entries[1].Subject)
	require.Len(t, entries[1].Messages, 2)
	require.Equal(t, "msg1 2", entries[1].Messages[0])
	require.Equal(t, "msg2", entries[1].Messages[1])
}
