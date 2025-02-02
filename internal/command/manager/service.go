package manager

import "path/filepath"

// NewPluginService to set up the plugin service
func NewPluginService(placeholder, storagePath string) IManager {
	return &SManager{
		placeholder: placeholder,
		storagePath: storagePath,
		pluginsPath: filepath.Join(storagePath, placeholder),
	}
}
