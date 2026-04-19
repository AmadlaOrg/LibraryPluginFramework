package output

import (
	"bytes"
	"testing"

	"github.com/AmadlaOrg/LibraryPluginFramework/protocol"
	"github.com/stretchr/testify/assert"
)

func TestWrite_JSON(t *testing.T) {
	var buf bytes.Buffer
	out := NewWithWriter(protocol.FormatJSON, &buf)

	data := map[string]any{"status": "pass", "count": 42}
	err := out.Write(data)

	assert.NoError(t, err)
	assert.Contains(t, buf.String(), `"status": "pass"`)
	assert.Contains(t, buf.String(), `"count": 42`)
}

func TestWrite_YAML(t *testing.T) {
	var buf bytes.Buffer
	out := NewWithWriter(protocol.FormatYAML, &buf)

	data := map[string]any{"status": "pass"}
	err := out.Write(data)

	assert.NoError(t, err)
	assert.Contains(t, buf.String(), "status: pass")
}

func TestWrite_Table(t *testing.T) {
	var buf bytes.Buffer
	out := NewWithWriter(protocol.FormatTable, &buf)

	data := map[string]any{"name": "test", "version": "1.0.0"}
	err := out.Write(data)

	assert.NoError(t, err)
	// tablewriter uppercases headers
	assert.Contains(t, buf.String(), "KEY")
	assert.Contains(t, buf.String(), "VALUE")
}

func TestWriteTable(t *testing.T) {
	var buf bytes.Buffer
	out := NewWithWriter(protocol.FormatTable, &buf)

	headers := []string{"Plugin", "Version", "Status"}
	rows := [][]string{
		{"doorman-vault", "1.0.0", "Active"},
		{"doorman-aws", "0.1.0", "Stub"},
	}
	err := out.WriteTable(headers, rows)

	assert.NoError(t, err)
	assert.Contains(t, buf.String(), "doorman-vault")
	assert.Contains(t, buf.String(), "doorman-aws")
	assert.Contains(t, buf.String(), "PLUGIN")
}

func TestWriteTable_JSON(t *testing.T) {
	var buf bytes.Buffer
	out := NewWithWriter(protocol.FormatJSON, &buf)

	headers := []string{"Plugin", "Version"}
	rows := [][]string{
		{"doorman-vault", "1.0.0"},
	}
	err := out.WriteTable(headers, rows)

	assert.NoError(t, err)
	assert.Contains(t, buf.String(), `"Plugin": "doorman-vault"`)
}

func TestWriteTable_YAML(t *testing.T) {
	var buf bytes.Buffer
	out := NewWithWriter(protocol.FormatYAML, &buf)

	headers := []string{"Plugin", "Version"}
	rows := [][]string{
		{"doorman-vault", "1.0.0"},
	}
	err := out.WriteTable(headers, rows)

	assert.NoError(t, err)
	assert.Contains(t, buf.String(), "Plugin: doorman-vault")
}

func TestParseFormatFlag(t *testing.T) {
	tests := []struct {
		input    string
		expected protocol.Format
	}{
		{"json", protocol.FormatJSON},
		{"yaml", protocol.FormatYAML},
		{"table", protocol.FormatTable},
		{"", protocol.FormatTable},
		{"invalid", protocol.FormatTable},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, ParseFormatFlag(tt.input))
		})
	}
}

func TestNew(t *testing.T) {
	out := New(protocol.FormatJSON)
	assert.NotNil(t, out)
}
