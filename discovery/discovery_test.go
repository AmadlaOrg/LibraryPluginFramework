package discovery

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestPlugins(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	// Create fake plugin binaries
	plugins := []string{"doorman-vault", "doorman-aws", "weaver-jinja", "unrelated-tool"}
	for _, name := range plugins {
		path := filepath.Join(dir, name)
		err := os.WriteFile(path, []byte("#!/bin/sh\n"), 0755)
		assert.NoError(t, err)
	}

	// Create a non-executable file that matches the prefix
	err := os.WriteFile(filepath.Join(dir, "doorman-notexec"), []byte("data"), 0644)
	assert.NoError(t, err)

	// Create a directory that matches the prefix
	err = os.Mkdir(filepath.Join(dir, "doorman-dir"), 0755)
	assert.NoError(t, err)

	return dir
}

func TestDiscover(t *testing.T) {
	dir := setupTestPlugins(t)

	// Override PATH lookup
	origLookPath := lookPath
	lookPath = func(key string) string {
		if key == "PATH" {
			return dir
		}
		return ""
	}
	defer func() { lookPath = origLookPath }()

	svc := New("doorman")
	plugins, err := svc.Discover()

	assert.NoError(t, err)
	assert.Len(t, plugins, 2) // doorman-vault, doorman-aws (not doorman-notexec, not doorman-dir)

	names := make([]string, len(plugins))
	for i, p := range plugins {
		names[i] = filepath.Base(p)
	}
	assert.Contains(t, names, "doorman-vault")
	assert.Contains(t, names, "doorman-aws")
	assert.NotContains(t, names, "doorman-notexec")
	assert.NotContains(t, names, "doorman-dir")
	assert.NotContains(t, names, "weaver-jinja")
}

func TestDiscoverNames(t *testing.T) {
	dir := setupTestPlugins(t)

	origLookPath := lookPath
	lookPath = func(key string) string {
		if key == "PATH" {
			return dir
		}
		return ""
	}
	defer func() { lookPath = origLookPath }()

	svc := New("doorman")
	names, err := svc.DiscoverNames()

	assert.NoError(t, err)
	assert.Len(t, names, 2)
	assert.Contains(t, names, "doorman-vault")
	assert.Contains(t, names, "doorman-aws")
}

func TestDiscover_EmptyPATH(t *testing.T) {
	origLookPath := lookPath
	lookPath = func(key string) string { return "" }
	defer func() { lookPath = origLookPath }()

	svc := New("doorman")
	plugins, err := svc.Discover()

	assert.NoError(t, err)
	assert.Nil(t, plugins)
}

func TestDiscover_NonexistentDir(t *testing.T) {
	origLookPath := lookPath
	lookPath = func(key string) string {
		if key == "PATH" {
			return "/nonexistent/path"
		}
		return ""
	}
	defer func() { lookPath = origLookPath }()

	svc := New("doorman")
	plugins, err := svc.Discover()

	assert.NoError(t, err)
	assert.Empty(t, plugins)
}

func TestDiscover_DuplicatesFirstWins(t *testing.T) {
	dir1 := t.TempDir()
	dir2 := t.TempDir()

	// Same plugin name in both dirs
	os.WriteFile(filepath.Join(dir1, "doorman-vault"), []byte("first"), 0755)
	os.WriteFile(filepath.Join(dir2, "doorman-vault"), []byte("second"), 0755)

	origLookPath := lookPath
	lookPath = func(key string) string {
		if key == "PATH" {
			return dir1 + string(os.PathListSeparator) + dir2
		}
		return ""
	}
	defer func() { lookPath = origLookPath }()

	svc := New("doorman")
	plugins, err := svc.Discover()

	assert.NoError(t, err)
	assert.Len(t, plugins, 1)
	assert.Equal(t, filepath.Join(dir1, "doorman-vault"), plugins[0]) // First on PATH wins
}

func TestIsExecutable(t *testing.T) {
	dir := t.TempDir()

	// Executable file
	execPath := filepath.Join(dir, "exec")
	os.WriteFile(execPath, []byte("#!/bin/sh"), 0755)

	// Non-executable file
	noexecPath := filepath.Join(dir, "noexec")
	os.WriteFile(noexecPath, []byte("data"), 0644)

	entries, _ := os.ReadDir(dir)
	entryMap := make(map[string]fs.DirEntry)
	for _, e := range entries {
		entryMap[e.Name()] = e
	}

	assert.True(t, isExecutable(entryMap["exec"]))
	assert.False(t, isExecutable(entryMap["noexec"]))
}
