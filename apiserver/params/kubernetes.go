// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package params

import (
	"github.com/juju/juju/constraints"
)

// KubernetesProvisioningInfo holds unit provisioning info.
type KubernetesProvisioningInfo struct {
	PodSpec               string                       `json:"pod-spec"`
	Constraints           constraints.Value            `json:"constraints"`
	Filesystems           []FilesystemParams           `json:"filesystems,omitempty"`
	FilesystemAttachments []FilesystemAttachmentParams `json:"filesystem-attachments,omitempty"`
	Volumes               []VolumeParams               `json:"volumes,omitempty"`
	VolumeAttachments     []VolumeAttachmentParams     `json:"volume-attachments,omitempty"`
	Tags                  map[string]string            `json:"tags,omitempty"`
}

// KubernetesProvisioningInfoResult holds unit provisioning info or an error.
type KubernetesProvisioningInfoResult struct {
	Error  *Error                      `json:"error,omitempty"`
	Result *KubernetesProvisioningInfo `json:"result"`
}

// KubernetesProvisioningInfoResults holds multiple provisioning info results.
type KubernetesProvisioningInfoResults struct {
	Results []KubernetesProvisioningInfoResult `json:"results"`
}
