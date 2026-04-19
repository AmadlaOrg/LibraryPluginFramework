package protocol

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

// Format represents a data format.
type Format string

const (
	FormatJSON    Format = "json"
	FormatYAML    Format = "yaml"
	FormatTable   Format = "table"
	FormatUnknown Format = "unknown"
)

// DetectFormat determines whether the given data is JSON or YAML.
// If the first non-whitespace character is '{' or '[', it's JSON; otherwise YAML.
func DetectFormat(data []byte) Format {
	trimmed := strings.TrimSpace(string(data))
	if len(trimmed) == 0 {
		return FormatUnknown
	}
	if trimmed[0] == '{' || trimmed[0] == '[' {
		return FormatJSON
	}
	return FormatYAML
}

// ParseEntity parses entity data (JSON or YAML) into a generic map.
// The format is auto-detected.
func ParseEntity(data []byte) (map[string]any, error) {
	format := DetectFormat(data)

	var result map[string]any
	switch format {
	case FormatJSON:
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, err
		}
	default:
		if err := yaml.Unmarshal(data, &result); err != nil {
			return nil, err
		}
	}
	return result, nil
}
