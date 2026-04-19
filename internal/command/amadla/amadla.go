package amadla

import "github.com/spf13/cobra"

type Amadla interface {
	Supported() map[string]any
	Cmd() *cobra.Command
	AmadlaCmd() *cobra.Command
}

type amadlaImpl struct {
	supportedApplications map[string]string
	supportedEntities     map[string]string
	amadlaCmd             *cobra.Command
}

// Cmd
func (s *amadlaImpl) Cmd() *cobra.Command {
	return s.amadlaCmd
}

// Supported
func (s *amadlaImpl) Supported() map[string]any {
	// Set default values if needed
	s.processSupportedApplications()

	return map[string]any{
		"applications": s.supportedApplications,
		"entities":     s.supportedEntities,
	}
}

// AmadlaCmd
func (s *amadlaImpl) AmadlaCmd() *cobra.Command {
	return s.amadlaCmd
}

// processSupportedApplications
func (s *amadlaImpl) processSupportedApplications() {
	var (
		heryPresent  bool
		judgePresent bool
	)

	for appName, appVersion := range s.supportedApplications {
		if appName == "hery" && appVersion != "" {
			heryPresent = true
		} else if appName == "judge" && appVersion != "" {
			judgePresent = true
		}
	}

	if !heryPresent {
		s.supportedApplications["hery"] = "^0"
	}
	if !judgePresent {
		s.supportedApplications["judge"] = "^0"
	}
}
