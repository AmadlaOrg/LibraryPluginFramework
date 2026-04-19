package manager

import "path/filepath"

// NewManagerService to set up the plugin manager service
func New(placeholder, storagePath string) Manager {
	return &managerImpl{
		placeholder: placeholder,
		storagePath: storagePath,
		pluginsPath: filepath.Join(storagePath, placeholder),
	}
}
