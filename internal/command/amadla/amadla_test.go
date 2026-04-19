package amadla

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCmd(t *testing.T) {
	cmd := &cobra.Command{Use: "amadla"}
	s := &amadlaImpl{amadlaCmd: cmd}
	assert.Equal(t, cmd, s.Cmd())
}

func TestAmadlaCmd(t *testing.T) {
	cmd := &cobra.Command{Use: "amadla"}
	s := &amadlaImpl{amadlaCmd: cmd}
	assert.Equal(t, cmd, s.AmadlaCmd())
}

func TestSupported_DefaultApplications(t *testing.T) {
	s := &amadlaImpl{
		supportedApplications: map[string]string{},
		supportedEntities:     map[string]string{},
	}

	result := s.Supported()
	apps := result["applications"].(map[string]string)
	assert.Equal(t, "^0", apps["hery"])
	assert.Equal(t, "^0", apps["judge"])
}

func TestSupported_ExistingApplications(t *testing.T) {
	s := &amadlaImpl{
		supportedApplications: map[string]string{
			"hery":  "^1",
			"judge": "^2",
		},
		supportedEntities: map[string]string{},
	}

	result := s.Supported()
	apps := result["applications"].(map[string]string)
	assert.Equal(t, "^1", apps["hery"])
	assert.Equal(t, "^2", apps["judge"])
}

func TestProcessSupportedApplications_OnlyAddsDefaults(t *testing.T) {
	s := &amadlaImpl{
		supportedApplications: map[string]string{
			"custom": "v1",
		},
		supportedEntities: map[string]string{},
	}

	s.processSupportedApplications()
	assert.Equal(t, "^0", s.supportedApplications["hery"])
	assert.Equal(t, "^0", s.supportedApplications["judge"])
	assert.Equal(t, "v1", s.supportedApplications["custom"])
}
