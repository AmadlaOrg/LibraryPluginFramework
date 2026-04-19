package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Format
	}{
		{"JSON object", `{"key": "value"}`, FormatJSON},
		{"JSON array", `[1, 2, 3]`, FormatJSON},
		{"JSON with whitespace", `  { "key": "value" }`, FormatJSON},
		{"YAML", "_type: test\n_body:\n  key: value", FormatYAML},
		{"YAML with leading whitespace", "  _type: test", FormatYAML},
		{"empty input", "", FormatUnknown},
		{"whitespace only", "   ", FormatUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, DetectFormat([]byte(tt.input)))
		})
	}
}

func TestParseEntity(t *testing.T) {
	t.Run("JSON input", func(t *testing.T) {
		data := []byte(`{"_type": "amadla.org/entity/application@v1.0.0", "_body": {"name": "myapp"}}`)
		result, err := ParseEntity(data)
		assert.NoError(t, err)
		assert.Equal(t, "amadla.org/entity/application@v1.0.0", result["_type"])
	})

	t.Run("YAML input", func(t *testing.T) {
		data := []byte("_type: amadla.org/entity/application@v1.0.0\n_body:\n  name: myapp")
		result, err := ParseEntity(data)
		assert.NoError(t, err)
		assert.Equal(t, "amadla.org/entity/application@v1.0.0", result["_type"])
	})

	t.Run("invalid JSON", func(t *testing.T) {
		data := []byte(`{invalid`)
		_, err := ParseEntity(data)
		assert.Error(t, err)
	})

	t.Run("invalid YAML", func(t *testing.T) {
		data := []byte(":\n  :\n    - :\n      :")
		_, err := ParseEntity(data)
		// YAML is very permissive, so this may or may not error
		// The point is ParseEntity doesn't panic
		_ = err
	})
}
