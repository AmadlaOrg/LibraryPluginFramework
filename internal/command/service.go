package command

import (
	"github.com/AmadlaOrg/LibraryPluginFramework/internal/command/amadla"
	"github.com/spf13/cobra"
)

// NewCommandService
func NewCommandService(cmd *cobra.Command,
	supportedApplications, supportedEntities map[string]string,
) {
	amadlaService := amadla.New(supportedApplications, supportedEntities)

	cmd.AddCommand(amadlaService.AmadlaCmd())
}
