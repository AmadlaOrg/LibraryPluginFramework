package discovery

import (
	"os"
	"path/filepath"
	"strings"
)

// Discovery scans the system PATH for plugin binaries matching a naming convention.
type Discovery interface {
	// Discover finds all plugin binaries on PATH matching the prefix "<toolName>-".
	// Returns a list of absolute paths to the discovered plugin binaries.
	Discover() ([]string, error)

	// DiscoverNames finds all plugin binaries on PATH and returns just their base names.
	DiscoverNames() ([]string, error)
}

// discoveryImpl implements Discovery.
type discoveryImpl struct {
	toolName string
}

// For testing: allow overriding PATH lookup
var lookPath = os.Getenv

// Discover finds all plugin binaries on PATH matching "<toolName>-*".
func (s *discoveryImpl) Discover() ([]string, error) {
	prefix := s.toolName + "-"
	pathEnv := lookPath("PATH")
	if pathEnv == "" {
		return nil, nil
	}

	seen := make(map[string]bool)
	var plugins []string

	dirs := filepath.SplitList(pathEnv)
	for _, dir := range dirs {
		entries, err := readDir(dir)
		if err != nil {
			// Skip directories we can't read
			continue
		}
		for _, entry := range entries {
			name := entry.Name()
			if !strings.HasPrefix(name, prefix) {
				continue
			}
			if seen[name] {
				// First match on PATH wins
				continue
			}
			fullPath := filepath.Join(dir, name)
			if isExecutable(entry) {
				seen[name] = true
				plugins = append(plugins, fullPath)
			}
		}
	}

	return plugins, nil
}

// DiscoverNames finds all plugin binaries and returns just the base names.
func (s *discoveryImpl) DiscoverNames() ([]string, error) {
	paths, err := s.Discover()
	if err != nil {
		return nil, err
	}
	names := make([]string, len(paths))
	for i, p := range paths {
		names[i] = filepath.Base(p)
	}
	return names, nil
}

// readDir wraps os.ReadDir for testability.
var readDir = os.ReadDir

// isExecutable checks if a directory entry is an executable file.
func isExecutable(entry os.DirEntry) bool {
	info, err := entry.Info()
	if err != nil {
		return false
	}
	// Must be a regular file (not a directory or symlink to a directory)
	if !info.Mode().IsRegular() {
		return false
	}
	// Check executable bit (owner, group, or other)
	return info.Mode().Perm()&0111 != 0
}
