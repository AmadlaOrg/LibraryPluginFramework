package amadla

import (
	"github.com/AmadlaOrg/LibraryJudgeFramework/display"
	"github.com/spf13/cobra"
)

// NewAmadlaService to set up the amadla service
func New(supportedApplications, supportedEntities map[string]string) Amadla {
	amadla := &amadlaImpl{
		supportedApplications: supportedApplications,
		supportedEntities:     supportedEntities,
		amadlaCmd: &cobra.Command{
			Use:   "amadla",
			Short: "Amadla supported applications and entities",
			Long:  `Displays the supported applications and entities. Use -o json|yaml|table to control output format.`,
		},
	}

	amadla.amadlaCmd.Run = func(cmd *cobra.Command, args []string) {
		if err := display.New(cmd, amadla.Supported()).
			SetTableHeaders([]string{"Category", "Supported", "Version Supported"}).
			Display(); err != nil {
			cmd.PrintErrln(err)
		}
	}

	return amadla
}
