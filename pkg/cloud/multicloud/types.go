package multicloud

// VirtualMachine represents a virtual machine resource.
type VirtualMachine struct {
	InstanceID   string
	InstanceName string
	CPUCores     int
	MemoryGB     int
	HourlyCost   float64
}

// ID returns the unique identifier of the virtual machine.
func (vm *VirtualMachine) ID() string {
	return vm.InstanceID
}

// Name returns the name of the virtual machine.
func (vm *VirtualMachine) Name() string {
	return vm.InstanceName
}

// ResourceType returns the type of the resource.
func (vm *VirtualMachine) ResourceType() string {
	return "VirtualMachine"
}

// Cost returns the hourly cost of the virtual machine.
func (vm *VirtualMachine) Cost() float64 {
	return vm.HourlyCost
}

// Storage represents a storage resource.
type Storage struct {
	StorageID   string
	StorageName string
	SizeGB      int
	HourlyCost  float64
}

// ID returns the unique identifier of the storage resource.
func (s *Storage) ID() string {
	return s.StorageID
}

// Name returns the name of the storage resource.
func (s *Storage) Name() string {
	return s.StorageName
}

// ResourceType returns the type of the resource.
func (s *Storage) ResourceType() string {
	return "Storage"
}

// Cost returns the hourly cost of the storage resource.
func (s *Storage) Cost() float64 {
	return s.HourlyCost
}
