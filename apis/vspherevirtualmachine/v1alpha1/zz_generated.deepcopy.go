//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CdromObservation) DeepCopyInto(out *CdromObservation) {
	*out = *in
	if in.DeviceAddress != nil {
		in, out := &in.DeviceAddress, &out.DeviceAddress
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CdromObservation.
func (in *CdromObservation) DeepCopy() *CdromObservation {
	if in == nil {
		return nil
	}
	out := new(CdromObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CdromParameters) DeepCopyInto(out *CdromParameters) {
	*out = *in
	if in.ClientDevice != nil {
		in, out := &in.ClientDevice, &out.ClientDevice
		*out = new(bool)
		**out = **in
	}
	if in.DatastoreID != nil {
		in, out := &in.DatastoreID, &out.DatastoreID
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CdromParameters.
func (in *CdromParameters) DeepCopy() *CdromParameters {
	if in == nil {
		return nil
	}
	out := new(CdromParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneObservation) DeepCopyInto(out *CloneObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneObservation.
func (in *CloneObservation) DeepCopy() *CloneObservation {
	if in == nil {
		return nil
	}
	out := new(CloneObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneParameters) DeepCopyInto(out *CloneParameters) {
	*out = *in
	if in.Customize != nil {
		in, out := &in.Customize, &out.Customize
		*out = make([]CustomizeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LinkedClone != nil {
		in, out := &in.LinkedClone, &out.LinkedClone
		*out = new(bool)
		**out = **in
	}
	if in.OvfNetworkMap != nil {
		in, out := &in.OvfNetworkMap, &out.OvfNetworkMap
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.OvfStorageMap != nil {
		in, out := &in.OvfStorageMap, &out.OvfStorageMap
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TemplateUUID != nil {
		in, out := &in.TemplateUUID, &out.TemplateUUID
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneParameters.
func (in *CloneParameters) DeepCopy() *CloneParameters {
	if in == nil {
		return nil
	}
	out := new(CloneParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizeObservation) DeepCopyInto(out *CustomizeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizeObservation.
func (in *CustomizeObservation) DeepCopy() *CustomizeObservation {
	if in == nil {
		return nil
	}
	out := new(CustomizeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizeParameters) DeepCopyInto(out *CustomizeParameters) {
	*out = *in
	if in.DNSServerList != nil {
		in, out := &in.DNSServerList, &out.DNSServerList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DNSSuffixList != nil {
		in, out := &in.DNSSuffixList, &out.DNSSuffixList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IPv4Gateway != nil {
		in, out := &in.IPv4Gateway, &out.IPv4Gateway
		*out = new(string)
		**out = **in
	}
	if in.IPv6Gateway != nil {
		in, out := &in.IPv6Gateway, &out.IPv6Gateway
		*out = new(string)
		**out = **in
	}
	if in.LinuxOptions != nil {
		in, out := &in.LinuxOptions, &out.LinuxOptions
		*out = make([]LinuxOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]NetworkInterfaceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
	if in.WindowsOptions != nil {
		in, out := &in.WindowsOptions, &out.WindowsOptions
		*out = make([]WindowsOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WindowsSysprepTextSecretRef != nil {
		in, out := &in.WindowsSysprepTextSecretRef, &out.WindowsSysprepTextSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizeParameters.
func (in *CustomizeParameters) DeepCopy() *CustomizeParameters {
	if in == nil {
		return nil
	}
	out := new(CustomizeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskObservation) DeepCopyInto(out *DiskObservation) {
	*out = *in
	if in.DeviceAddress != nil {
		in, out := &in.DeviceAddress, &out.DeviceAddress
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(float64)
		**out = **in
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskObservation.
func (in *DiskObservation) DeepCopy() *DiskObservation {
	if in == nil {
		return nil
	}
	out := new(DiskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskParameters) DeepCopyInto(out *DiskParameters) {
	*out = *in
	if in.Attach != nil {
		in, out := &in.Attach, &out.Attach
		*out = new(bool)
		**out = **in
	}
	if in.ControllerType != nil {
		in, out := &in.ControllerType, &out.ControllerType
		*out = new(string)
		**out = **in
	}
	if in.DatastoreID != nil {
		in, out := &in.DatastoreID, &out.DatastoreID
		*out = new(string)
		**out = **in
	}
	if in.DiskMode != nil {
		in, out := &in.DiskMode, &out.DiskMode
		*out = new(string)
		**out = **in
	}
	if in.DiskSharing != nil {
		in, out := &in.DiskSharing, &out.DiskSharing
		*out = new(string)
		**out = **in
	}
	if in.EagerlyScrub != nil {
		in, out := &in.EagerlyScrub, &out.EagerlyScrub
		*out = new(bool)
		**out = **in
	}
	if in.IoLimit != nil {
		in, out := &in.IoLimit, &out.IoLimit
		*out = new(float64)
		**out = **in
	}
	if in.IoReservation != nil {
		in, out := &in.IoReservation, &out.IoReservation
		*out = new(float64)
		**out = **in
	}
	if in.IoShareCount != nil {
		in, out := &in.IoShareCount, &out.IoShareCount
		*out = new(float64)
		**out = **in
	}
	if in.IoShareLevel != nil {
		in, out := &in.IoShareLevel, &out.IoShareLevel
		*out = new(string)
		**out = **in
	}
	if in.KeepOnRemove != nil {
		in, out := &in.KeepOnRemove, &out.KeepOnRemove
		*out = new(bool)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.StoragePolicyID != nil {
		in, out := &in.StoragePolicyID, &out.StoragePolicyID
		*out = new(string)
		**out = **in
	}
	if in.ThinProvisioned != nil {
		in, out := &in.ThinProvisioned, &out.ThinProvisioned
		*out = new(bool)
		**out = **in
	}
	if in.UnitNumber != nil {
		in, out := &in.UnitNumber, &out.UnitNumber
		*out = new(float64)
		**out = **in
	}
	if in.WriteThrough != nil {
		in, out := &in.WriteThrough, &out.WriteThrough
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskParameters.
func (in *DiskParameters) DeepCopy() *DiskParameters {
	if in == nil {
		return nil
	}
	out := new(DiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinuxOptionsObservation) DeepCopyInto(out *LinuxOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinuxOptionsObservation.
func (in *LinuxOptionsObservation) DeepCopy() *LinuxOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(LinuxOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinuxOptionsParameters) DeepCopyInto(out *LinuxOptionsParameters) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.HwClockUtc != nil {
		in, out := &in.HwClockUtc, &out.HwClockUtc
		*out = new(bool)
		**out = **in
	}
	if in.ScriptTextSecretRef != nil {
		in, out := &in.ScriptTextSecretRef, &out.ScriptTextSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinuxOptionsParameters.
func (in *LinuxOptionsParameters) DeepCopy() *LinuxOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(LinuxOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Machine) DeepCopyInto(out *Machine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Machine.
func (in *Machine) DeepCopy() *Machine {
	if in == nil {
		return nil
	}
	out := new(Machine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Machine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineList) DeepCopyInto(out *MachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Machine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineList.
func (in *MachineList) DeepCopy() *MachineList {
	if in == nil {
		return nil
	}
	out := new(MachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineNetworkInterfaceObservation) DeepCopyInto(out *MachineNetworkInterfaceObservation) {
	*out = *in
	if in.DeviceAddress != nil {
		in, out := &in.DeviceAddress, &out.DeviceAddress
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineNetworkInterfaceObservation.
func (in *MachineNetworkInterfaceObservation) DeepCopy() *MachineNetworkInterfaceObservation {
	if in == nil {
		return nil
	}
	out := new(MachineNetworkInterfaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineNetworkInterfaceParameters) DeepCopyInto(out *MachineNetworkInterfaceParameters) {
	*out = *in
	if in.AdapterType != nil {
		in, out := &in.AdapterType, &out.AdapterType
		*out = new(string)
		**out = **in
	}
	if in.BandwidthLimit != nil {
		in, out := &in.BandwidthLimit, &out.BandwidthLimit
		*out = new(float64)
		**out = **in
	}
	if in.BandwidthReservation != nil {
		in, out := &in.BandwidthReservation, &out.BandwidthReservation
		*out = new(float64)
		**out = **in
	}
	if in.BandwidthShareCount != nil {
		in, out := &in.BandwidthShareCount, &out.BandwidthShareCount
		*out = new(float64)
		**out = **in
	}
	if in.BandwidthShareLevel != nil {
		in, out := &in.BandwidthShareLevel, &out.BandwidthShareLevel
		*out = new(string)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.OvfMapping != nil {
		in, out := &in.OvfMapping, &out.OvfMapping
		*out = new(string)
		**out = **in
	}
	if in.UseStaticMac != nil {
		in, out := &in.UseStaticMac, &out.UseStaticMac
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineNetworkInterfaceParameters.
func (in *MachineNetworkInterfaceParameters) DeepCopy() *MachineNetworkInterfaceParameters {
	if in == nil {
		return nil
	}
	out := new(MachineNetworkInterfaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineObservation) DeepCopyInto(out *MachineObservation) {
	*out = *in
	if in.ChangeVersion != nil {
		in, out := &in.ChangeVersion, &out.ChangeVersion
		*out = new(string)
		**out = **in
	}
	if in.DefaultIPAddress != nil {
		in, out := &in.DefaultIPAddress, &out.DefaultIPAddress
		*out = new(string)
		**out = **in
	}
	if in.GuestIPAddresses != nil {
		in, out := &in.GuestIPAddresses, &out.GuestIPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Imported != nil {
		in, out := &in.Imported, &out.Imported
		*out = new(bool)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.PowerState != nil {
		in, out := &in.PowerState, &out.PowerState
		*out = new(string)
		**out = **in
	}
	if in.RebootRequired != nil {
		in, out := &in.RebootRequired, &out.RebootRequired
		*out = new(bool)
		**out = **in
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
	if in.VappTransport != nil {
		in, out := &in.VappTransport, &out.VappTransport
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VmwareToolsStatus != nil {
		in, out := &in.VmwareToolsStatus, &out.VmwareToolsStatus
		*out = new(string)
		**out = **in
	}
	if in.VmxPath != nil {
		in, out := &in.VmxPath, &out.VmxPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineObservation.
func (in *MachineObservation) DeepCopy() *MachineObservation {
	if in == nil {
		return nil
	}
	out := new(MachineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineParameters) DeepCopyInto(out *MachineParameters) {
	*out = *in
	if in.AlternateGuestName != nil {
		in, out := &in.AlternateGuestName, &out.AlternateGuestName
		*out = new(string)
		**out = **in
	}
	if in.Annotation != nil {
		in, out := &in.Annotation, &out.Annotation
		*out = new(string)
		**out = **in
	}
	if in.BootDelay != nil {
		in, out := &in.BootDelay, &out.BootDelay
		*out = new(float64)
		**out = **in
	}
	if in.BootRetryDelay != nil {
		in, out := &in.BootRetryDelay, &out.BootRetryDelay
		*out = new(float64)
		**out = **in
	}
	if in.BootRetryEnabled != nil {
		in, out := &in.BootRetryEnabled, &out.BootRetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CPUHotAddEnabled != nil {
		in, out := &in.CPUHotAddEnabled, &out.CPUHotAddEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CPUHotRemoveEnabled != nil {
		in, out := &in.CPUHotRemoveEnabled, &out.CPUHotRemoveEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CPULimit != nil {
		in, out := &in.CPULimit, &out.CPULimit
		*out = new(float64)
		**out = **in
	}
	if in.CPUPerformanceCountersEnabled != nil {
		in, out := &in.CPUPerformanceCountersEnabled, &out.CPUPerformanceCountersEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CPUReservation != nil {
		in, out := &in.CPUReservation, &out.CPUReservation
		*out = new(float64)
		**out = **in
	}
	if in.CPUShareCount != nil {
		in, out := &in.CPUShareCount, &out.CPUShareCount
		*out = new(float64)
		**out = **in
	}
	if in.CPUShareLevel != nil {
		in, out := &in.CPUShareLevel, &out.CPUShareLevel
		*out = new(string)
		**out = **in
	}
	if in.Cdrom != nil {
		in, out := &in.Cdrom, &out.Cdrom
		*out = make([]CdromParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = make([]CloneParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.DatastoreID != nil {
		in, out := &in.DatastoreID, &out.DatastoreID
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = make([]DiskParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EfiSecureBootEnabled != nil {
		in, out := &in.EfiSecureBootEnabled, &out.EfiSecureBootEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EnableDiskUUID != nil {
		in, out := &in.EnableDiskUUID, &out.EnableDiskUUID
		*out = new(bool)
		**out = **in
	}
	if in.EnableLogging != nil {
		in, out := &in.EnableLogging, &out.EnableLogging
		*out = new(bool)
		**out = **in
	}
	if in.EptRviMode != nil {
		in, out := &in.EptRviMode, &out.EptRviMode
		*out = new(string)
		**out = **in
	}
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Firmware != nil {
		in, out := &in.Firmware, &out.Firmware
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.ForcePowerOff != nil {
		in, out := &in.ForcePowerOff, &out.ForcePowerOff
		*out = new(bool)
		**out = **in
	}
	if in.GuestID != nil {
		in, out := &in.GuestID, &out.GuestID
		*out = new(string)
		**out = **in
	}
	if in.HardwareVersion != nil {
		in, out := &in.HardwareVersion, &out.HardwareVersion
		*out = new(float64)
		**out = **in
	}
	if in.HostSystemID != nil {
		in, out := &in.HostSystemID, &out.HostSystemID
		*out = new(string)
		**out = **in
	}
	if in.HvMode != nil {
		in, out := &in.HvMode, &out.HvMode
		*out = new(string)
		**out = **in
	}
	if in.IdeControllerCount != nil {
		in, out := &in.IdeControllerCount, &out.IdeControllerCount
		*out = new(float64)
		**out = **in
	}
	if in.IgnoredGuestIps != nil {
		in, out := &in.IgnoredGuestIps, &out.IgnoredGuestIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LatencySensitivity != nil {
		in, out := &in.LatencySensitivity, &out.LatencySensitivity
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.MemoryHotAddEnabled != nil {
		in, out := &in.MemoryHotAddEnabled, &out.MemoryHotAddEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(float64)
		**out = **in
	}
	if in.MemoryReservation != nil {
		in, out := &in.MemoryReservation, &out.MemoryReservation
		*out = new(float64)
		**out = **in
	}
	if in.MemoryShareCount != nil {
		in, out := &in.MemoryShareCount, &out.MemoryShareCount
		*out = new(float64)
		**out = **in
	}
	if in.MemoryShareLevel != nil {
		in, out := &in.MemoryShareLevel, &out.MemoryShareLevel
		*out = new(string)
		**out = **in
	}
	if in.MigrateWaitTimeout != nil {
		in, out := &in.MigrateWaitTimeout, &out.MigrateWaitTimeout
		*out = new(float64)
		**out = **in
	}
	if in.NestedHvEnabled != nil {
		in, out := &in.NestedHvEnabled, &out.NestedHvEnabled
		*out = new(bool)
		**out = **in
	}
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]MachineNetworkInterfaceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NumCoresPerSocket != nil {
		in, out := &in.NumCoresPerSocket, &out.NumCoresPerSocket
		*out = new(float64)
		**out = **in
	}
	if in.NumCpus != nil {
		in, out := &in.NumCpus, &out.NumCpus
		*out = new(float64)
		**out = **in
	}
	if in.OvfDeploy != nil {
		in, out := &in.OvfDeploy, &out.OvfDeploy
		*out = make([]OvfDeployParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PciDeviceID != nil {
		in, out := &in.PciDeviceID, &out.PciDeviceID
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PoweronTimeout != nil {
		in, out := &in.PoweronTimeout, &out.PoweronTimeout
		*out = new(float64)
		**out = **in
	}
	if in.ReplaceTrigger != nil {
		in, out := &in.ReplaceTrigger, &out.ReplaceTrigger
		*out = new(string)
		**out = **in
	}
	if in.ResourcePoolID != nil {
		in, out := &in.ResourcePoolID, &out.ResourcePoolID
		*out = new(string)
		**out = **in
	}
	if in.RunToolsScriptsAfterPowerOn != nil {
		in, out := &in.RunToolsScriptsAfterPowerOn, &out.RunToolsScriptsAfterPowerOn
		*out = new(bool)
		**out = **in
	}
	if in.RunToolsScriptsAfterResume != nil {
		in, out := &in.RunToolsScriptsAfterResume, &out.RunToolsScriptsAfterResume
		*out = new(bool)
		**out = **in
	}
	if in.RunToolsScriptsBeforeGuestReboot != nil {
		in, out := &in.RunToolsScriptsBeforeGuestReboot, &out.RunToolsScriptsBeforeGuestReboot
		*out = new(bool)
		**out = **in
	}
	if in.RunToolsScriptsBeforeGuestShutdown != nil {
		in, out := &in.RunToolsScriptsBeforeGuestShutdown, &out.RunToolsScriptsBeforeGuestShutdown
		*out = new(bool)
		**out = **in
	}
	if in.RunToolsScriptsBeforeGuestStandby != nil {
		in, out := &in.RunToolsScriptsBeforeGuestStandby, &out.RunToolsScriptsBeforeGuestStandby
		*out = new(bool)
		**out = **in
	}
	if in.SataControllerCount != nil {
		in, out := &in.SataControllerCount, &out.SataControllerCount
		*out = new(float64)
		**out = **in
	}
	if in.ScsiBusSharing != nil {
		in, out := &in.ScsiBusSharing, &out.ScsiBusSharing
		*out = new(string)
		**out = **in
	}
	if in.ScsiControllerCount != nil {
		in, out := &in.ScsiControllerCount, &out.ScsiControllerCount
		*out = new(float64)
		**out = **in
	}
	if in.ScsiType != nil {
		in, out := &in.ScsiType, &out.ScsiType
		*out = new(string)
		**out = **in
	}
	if in.ShutdownWaitTimeout != nil {
		in, out := &in.ShutdownWaitTimeout, &out.ShutdownWaitTimeout
		*out = new(float64)
		**out = **in
	}
	if in.StoragePolicyID != nil {
		in, out := &in.StoragePolicyID, &out.StoragePolicyID
		*out = new(string)
		**out = **in
	}
	if in.SwapPlacementPolicy != nil {
		in, out := &in.SwapPlacementPolicy, &out.SwapPlacementPolicy
		*out = new(string)
		**out = **in
	}
	if in.SyncTimeWithHost != nil {
		in, out := &in.SyncTimeWithHost, &out.SyncTimeWithHost
		*out = new(bool)
		**out = **in
	}
	if in.SyncTimeWithHostPeriodically != nil {
		in, out := &in.SyncTimeWithHostPeriodically, &out.SyncTimeWithHostPeriodically
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ToolsUpgradePolicy != nil {
		in, out := &in.ToolsUpgradePolicy, &out.ToolsUpgradePolicy
		*out = new(string)
		**out = **in
	}
	if in.Vapp != nil {
		in, out := &in.Vapp, &out.Vapp
		*out = make([]VappParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VbsEnabled != nil {
		in, out := &in.VbsEnabled, &out.VbsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.VvtdEnabled != nil {
		in, out := &in.VvtdEnabled, &out.VvtdEnabled
		*out = new(bool)
		**out = **in
	}
	if in.WaitForGuestIPTimeout != nil {
		in, out := &in.WaitForGuestIPTimeout, &out.WaitForGuestIPTimeout
		*out = new(float64)
		**out = **in
	}
	if in.WaitForGuestNetRoutable != nil {
		in, out := &in.WaitForGuestNetRoutable, &out.WaitForGuestNetRoutable
		*out = new(bool)
		**out = **in
	}
	if in.WaitForGuestNetTimeout != nil {
		in, out := &in.WaitForGuestNetTimeout, &out.WaitForGuestNetTimeout
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineParameters.
func (in *MachineParameters) DeepCopy() *MachineParameters {
	if in == nil {
		return nil
	}
	out := new(MachineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpec) DeepCopyInto(out *MachineSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpec.
func (in *MachineSpec) DeepCopy() *MachineSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineStatus) DeepCopyInto(out *MachineStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineStatus.
func (in *MachineStatus) DeepCopy() *MachineStatus {
	if in == nil {
		return nil
	}
	out := new(MachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceObservation) DeepCopyInto(out *NetworkInterfaceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceObservation.
func (in *NetworkInterfaceObservation) DeepCopy() *NetworkInterfaceObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceParameters) DeepCopyInto(out *NetworkInterfaceParameters) {
	*out = *in
	if in.DNSDomain != nil {
		in, out := &in.DNSDomain, &out.DNSDomain
		*out = new(string)
		**out = **in
	}
	if in.DNSServerList != nil {
		in, out := &in.DNSServerList, &out.DNSServerList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IPv4Address != nil {
		in, out := &in.IPv4Address, &out.IPv4Address
		*out = new(string)
		**out = **in
	}
	if in.IPv4Netmask != nil {
		in, out := &in.IPv4Netmask, &out.IPv4Netmask
		*out = new(float64)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.IPv6Netmask != nil {
		in, out := &in.IPv6Netmask, &out.IPv6Netmask
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceParameters.
func (in *NetworkInterfaceParameters) DeepCopy() *NetworkInterfaceParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvfDeployObservation) DeepCopyInto(out *OvfDeployObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvfDeployObservation.
func (in *OvfDeployObservation) DeepCopy() *OvfDeployObservation {
	if in == nil {
		return nil
	}
	out := new(OvfDeployObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvfDeployParameters) DeepCopyInto(out *OvfDeployParameters) {
	*out = *in
	if in.AllowUnverifiedSSLCert != nil {
		in, out := &in.AllowUnverifiedSSLCert, &out.AllowUnverifiedSSLCert
		*out = new(bool)
		**out = **in
	}
	if in.DeploymentOption != nil {
		in, out := &in.DeploymentOption, &out.DeploymentOption
		*out = new(string)
		**out = **in
	}
	if in.DiskProvisioning != nil {
		in, out := &in.DiskProvisioning, &out.DiskProvisioning
		*out = new(string)
		**out = **in
	}
	if in.EnableHiddenProperties != nil {
		in, out := &in.EnableHiddenProperties, &out.EnableHiddenProperties
		*out = new(bool)
		**out = **in
	}
	if in.IPAllocationPolicy != nil {
		in, out := &in.IPAllocationPolicy, &out.IPAllocationPolicy
		*out = new(string)
		**out = **in
	}
	if in.IPProtocol != nil {
		in, out := &in.IPProtocol, &out.IPProtocol
		*out = new(string)
		**out = **in
	}
	if in.LocalOvfPath != nil {
		in, out := &in.LocalOvfPath, &out.LocalOvfPath
		*out = new(string)
		**out = **in
	}
	if in.OvfNetworkMap != nil {
		in, out := &in.OvfNetworkMap, &out.OvfNetworkMap
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.RemoteOvfURL != nil {
		in, out := &in.RemoteOvfURL, &out.RemoteOvfURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvfDeployParameters.
func (in *OvfDeployParameters) DeepCopy() *OvfDeployParameters {
	if in == nil {
		return nil
	}
	out := new(OvfDeployParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VappObservation) DeepCopyInto(out *VappObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VappObservation.
func (in *VappObservation) DeepCopy() *VappObservation {
	if in == nil {
		return nil
	}
	out := new(VappObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VappParameters) DeepCopyInto(out *VappParameters) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VappParameters.
func (in *VappParameters) DeepCopy() *VappParameters {
	if in == nil {
		return nil
	}
	out := new(VappParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsOptionsObservation) DeepCopyInto(out *WindowsOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsOptionsObservation.
func (in *WindowsOptionsObservation) DeepCopy() *WindowsOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(WindowsOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsOptionsParameters) DeepCopyInto(out *WindowsOptionsParameters) {
	*out = *in
	if in.AdminPasswordSecretRef != nil {
		in, out := &in.AdminPasswordSecretRef, &out.AdminPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AutoLogon != nil {
		in, out := &in.AutoLogon, &out.AutoLogon
		*out = new(bool)
		**out = **in
	}
	if in.AutoLogonCount != nil {
		in, out := &in.AutoLogonCount, &out.AutoLogonCount
		*out = new(float64)
		**out = **in
	}
	if in.ComputerName != nil {
		in, out := &in.ComputerName, &out.ComputerName
		*out = new(string)
		**out = **in
	}
	if in.DomainAdminPasswordSecretRef != nil {
		in, out := &in.DomainAdminPasswordSecretRef, &out.DomainAdminPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.DomainAdminUser != nil {
		in, out := &in.DomainAdminUser, &out.DomainAdminUser
		*out = new(string)
		**out = **in
	}
	if in.FullName != nil {
		in, out := &in.FullName, &out.FullName
		*out = new(string)
		**out = **in
	}
	if in.JoinDomain != nil {
		in, out := &in.JoinDomain, &out.JoinDomain
		*out = new(string)
		**out = **in
	}
	if in.OrganizationName != nil {
		in, out := &in.OrganizationName, &out.OrganizationName
		*out = new(string)
		**out = **in
	}
	if in.ProductKeySecretRef != nil {
		in, out := &in.ProductKeySecretRef, &out.ProductKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.RunOnceCommandList != nil {
		in, out := &in.RunOnceCommandList, &out.RunOnceCommandList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(float64)
		**out = **in
	}
	if in.Workgroup != nil {
		in, out := &in.Workgroup, &out.Workgroup
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsOptionsParameters.
func (in *WindowsOptionsParameters) DeepCopy() *WindowsOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(WindowsOptionsParameters)
	in.DeepCopyInto(out)
	return out
}