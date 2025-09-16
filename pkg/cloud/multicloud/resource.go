package multicloud

// Resource represents a generic cloud resource.
type Resource interface {
	// ID returns the unique identifier of the resource.
	ID() string
	// Name returns the name of the resource.
	Name() string
	// ResourceType returns the type of the resource (e.g., "VirtualMachine", "StorageBucket").
	ResourceType() string
	// Cost returns the hourly cost of the resource.
	Cost() float64
}
