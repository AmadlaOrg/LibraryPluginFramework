package manager

import "path/filepath"

// NewManagerService to set up the plugin manager service
func NewManagerService(placeholder, storagePath string) IManager {
	return &SManager{
		placeholder: placeholder,
		storagePath: storagePath,
		pluginsPath: filepath.Join(storagePath, placeholder),
	}
}
