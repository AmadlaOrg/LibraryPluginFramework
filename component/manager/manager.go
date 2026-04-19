package manager

import (
	"github.com/AmadlaOrg/LibraryPluginFramework/internal/command/manager"
	"github.com/spf13/cobra"
)

// New attaches command for adding/removing/updating plugins (and for HERY)
//
// Params:
// - ⌨️ cmd - Is for attaching the command/commands to parent command
// - 🖼️ placeholder - (Default: `plugin` if it is set as empty) Is what will be the name of the command that will manage plugins, it is also used as the name of the directory where the plugins will be stored
// - 💾 storagePath - Is the absolute path to where the plugins directory (the name of the directory is taken from 🖼️ placeholder) is stored
//
// Example:
// - placeholder = plugin
// - storagePath = /home/user/.doorman/
// With those values the complete absolute path to where the plugins will be stored will be:
// - /home/user/.doorman/plugin/
func New(cmd *cobra.Command, placeholder, storagePath string, runPluginManager manager.RunPluginManager) manager.Manager {
	return manager.New(placeholder, storagePath)
}
