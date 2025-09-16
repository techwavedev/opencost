package multicloud

// CloudCost represents the total cost of all resources for a given cloud provider.
type CloudCost struct {
	Provider  string
	Resources []Resource
}
