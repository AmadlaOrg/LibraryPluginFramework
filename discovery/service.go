package discovery

// NewDiscoveryService creates a new discovery service for the given tool name.
// It will scan $PATH for binaries matching "<toolName>-*".
//
// Example:
//
//	svc := New("doorman")
//	plugins, _ := svc.Discover()  // finds doorman-vault, doorman-aws, etc.
func New(toolName string) Discovery {
	return &discoveryImpl{
		toolName: toolName,
	}
}
