package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/AmadlaOrg/LibraryPluginFramework/protocol"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

// Output formats and writes data to an output stream.
type Output interface {
	// Write formats the data according to the configured format and writes it to the output stream.
	Write(data any) error

	// WriteTable writes data as a formatted table with the given headers and rows.
	WriteTable(headers []string, rows [][]string) error
}

// outputImpl implements Output.
type outputImpl struct {
	format protocol.Format
	writer io.Writer
}

// For testing
var (
	jsonMarshalIndent = json.MarshalIndent
	yamlMarshal       = yaml.Marshal
)

// Write formats the data and writes it to the output stream.
func (s *outputImpl) Write(data any) error {
	switch s.format {
	case protocol.FormatJSON:
		return s.writeJSON(data)
	case protocol.FormatYAML:
		return s.writeYAML(data)
	case protocol.FormatTable:
		return s.writeKeyValue(data)
	default:
		return s.writeJSON(data)
	}
}

// WriteTable writes a formatted table.
func (s *outputImpl) WriteTable(headers []string, rows [][]string) error {
	if s.format == protocol.FormatJSON || s.format == protocol.FormatYAML {
		// Convert table to structured data for JSON/YAML output
		result := make([]map[string]string, len(rows))
		for i, row := range rows {
			entry := make(map[string]string)
			for j, header := range headers {
				if j < len(row) {
					entry[header] = row[j]
				}
			}
			result[i] = entry
		}
		return s.Write(result)
	}

	table := tablewriter.NewWriter(s.writer)
	table.Header(headers)
	table.Bulk(rows)
	table.Render()
	return nil
}

func (s *outputImpl) writeJSON(data any) error {
	bytes, err := jsonMarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}
	_, err = fmt.Fprintln(s.writer, string(bytes))
	return err
}

func (s *outputImpl) writeYAML(data any) error {
	bytes, err := yamlMarshal(data)
	if err != nil {
		return fmt.Errorf("error encoding YAML: %w", err)
	}
	_, err = fmt.Fprint(s.writer, string(bytes))
	return err
}

func (s *outputImpl) writeKeyValue(data any) error {
	// For table format with unstructured data, render as key-value pairs
	switch v := data.(type) {
	case map[string]any:
		table := tablewriter.NewWriter(s.writer)
		table.Header("Key", "Value")
		for key, val := range v {
			table.Append([]string{key, fmt.Sprintf("%v", val)})
		}
		table.Render()
	default:
		// Fall back to JSON for complex types in table mode
		return s.writeJSON(data)
	}
	return nil
}

// ParseFormatFlag parses the -o flag value into a Format.
func ParseFormatFlag(value string) protocol.Format {
	switch value {
	case "json":
		return protocol.FormatJSON
	case "yaml":
		return protocol.FormatYAML
	case "table", "":
		return protocol.FormatTable
	default:
		return protocol.FormatTable
	}
}

// NewOutputService creates an output service for the given format, writing to stdout.
func New(format protocol.Format) Output {
	return &outputImpl{
		format: format,
		writer: os.Stdout,
	}
}

// NewOutputServiceWithWriter creates an output service writing to a custom writer.
func NewWithWriter(format protocol.Format, writer io.Writer) Output {
	return &outputImpl{
		format: format,
		writer: writer,
	}
}
