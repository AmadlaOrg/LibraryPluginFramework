package protocol

import "encoding/json"

// Info represents the metadata returned by a plugin's `info` subcommand.
type Info struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Engine      string   `json:"engine,omitempty"`
	Supports    []string `json:"supports"`
	Description string   `json:"description"`
}

// ParseInfo parses JSON output from a plugin's `info` subcommand into an Info struct.
func ParseInfo(data []byte) (Info, error) {
	var info Info
	if err := json.Unmarshal(data, &info); err != nil {
		return Info{}, err
	}
	return info, nil
}

// SupportsEntity returns true if the plugin declares support for the given entity type URI.
// Matching is prefix-based: "amadla.org/entity/application@^v1" matches "amadla.org/entity/application@v1.2.3".
// For now, this does exact string matching. Semver range matching can be added later.
func (i Info) SupportsEntity(entityTypeURI string) bool {
	for _, s := range i.Supports {
		if s == entityTypeURI {
			return true
		}
	}
	return false
}
