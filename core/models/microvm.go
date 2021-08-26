package models

// MicroVM represents a microvm machine that is created via a provider.
type MicroVM struct {
	// ID is the identifier for the microvm.
	ID VMID `json:"id"`
	// Version is the version for the microvm definition.
	Version int `json:"version"`
	// Spec is the specification of the microvm.
	Spec MicroVMSpec `json:"spec"`
}

// MicroVMSpec represents the specification of a microvm machine.
type MicroVMSpec struct {
	// Kernel specifies the kernel and its argments to use.
	Kernel Kernel `json:"kernel"`
	// InitrdImage is an optional initial ramdisk to use.
	InitrdImage ContainerImage `json:"initrd_image,omitempty"`
	// VCPU specifies how many vcpu the machine will be allocated.
	VCPU int64 `json:"vcpu"`
	// MemoryInMb is the amount of memory in megabytes that the machine will be allocated.
	MemoryInMb int64 `json:"memory_inmb"`
	// NetworkInterfaces specifies the network interfaces attached to the machine.
	NetworkInterfaces []NetworkInterface `json:"network_interfaces"`
	// Volumes specifies the volumes to be attached to the the machine.
	Volumes []Volume `json:"volumes"`
}

// Kernel is the specification of the kernel and its arguments.
type Kernel struct {
	// Image is the container image to use for the kernel.
	Image ContainerImage `json:"image"`
	// Filename is the name of the kernel filename in the container.
	Filename string
	// CmdLine are the args to use for the kernel cmdline.
	CmdLine string `json:"cmdline,omitempty"`
}

// ContainerImage represents the address of a OCI image.
type ContainerImage string

// NetworkInterface represents a network interface for the microvm.
type NetworkInterface struct {
	// AllowMetadataRequests indicates that this interface can be used for metadata requests.
	// TODO: we may hide this within the firecracker plugin.
	AllowMetadataRequests bool `json:"allow_mmds,omitempty"`
	// GuestMAC allows the specifying of a specifi MAC address to use for the interface. If
	// not supplied a autogenerated MAC address will be used.
	GuestMAC string `json:"guest_mac,omitempty"`
	// HostDeviceName is the name of the network interface to use from the host. This will be
	// a tuntap or macvtap interface.
	HostDeviceName string `json:"host_device_name"`
	// GuestDeviceName is the name of the network interface to create in the microvm. If this
	// is not supplied than a device name will be assigned automatically.
	GuestDeviceName string `json:"guest_device_name,omitempty"`
	// TODO: add rate limiting.
	// TODO: add CNI.
}

// Volume represents a volume to be attached to a microvm machine.
type Volume struct {
	// ID is the uinique identifier of the volume.
	ID string `json:"id"`
	// IsRoot specifies that the volume is to be used as the root volume. A machine
	// must have a root volume.
	IsRoot bool `json:"is_root"`
	// IsReadOnly specifies that the volume is to be mounted readonly.
	IsReadOnly bool `json:"is_read_only,omitempty"`
	// MountPoint is the mount point for the volume in the microvm.
	MountPoint string `json:"mount_point"`
	// Source is where the volume will be sourced from.
	Source VolumeSource `json:"source"`
	// PartitionID is the uuid of the boot partition.
	PartitionID string `json:"partition_id,omitempty"`
	// Size is the size to resize this volume to.
	Size int32 `json:"size,omitempty"`
	// TODO: add rate limiting.
}

// VolumeSource is the source of a volume. Based loosely on the volumes in Kubernetes Pod specs.
type VolumeSource struct {
	// Container is used to specify a source of a volume as a OCI container.
	Container *ContainerVolumeSource `json:"container,omitempty"`
	// HostPath is used to specify a source of a volume as a file/device on the host.
	HostPath *HostPathVolumeSource `json:"host_path,omitempty"`
	// TODO: add CSI.
}

// ContainerDriveSource represents the details of a volume coming from a OCI image.
type ContainerVolumeSource struct {
	// Image is the OCI image to use.
	Image ContainerImage `json:"image"`
}

// HostPathVolumeSource represents the details of a volume coming from a file/device on the host.
type HostPathVolumeSource struct {
	// Path on the host of the file/device to use as a source for a volume.
	Path string
	// Type is the type of file/device on the host.
	Type HostPathType
}

// HostPathType is a type representing the different type of files/devices.
type HostPathType string

const (
	// HostPathRawFile represents a file on the host to use as a source for a volume. It should be a raw fs.
	HostPathRawFile HostPathType = "RawFile"
)
