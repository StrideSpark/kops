package api

import (
	"k8s.io/kubernetes/pkg/util/validation/field"
)

type DockerConfig struct {
	Bridge           *string `json:"bridge,omitempty" flag:"bridge"`
	LogLevel         *string `json:"logLevel,omitempty" flag:"log-level"`
	IPTables         *bool   `json:"ipTables,omitempty" flag:"iptables"`
	IPMasq           *bool   `json:"ipMasq,omitempty" flag:"ip-masq"`
	Storage          *string `json:"storage,omitempty" flag:"storage-driver"`
	InsecureRegistry *string `json:"insecureRegistry,omitempty" flag:"insecure-registry"`
	MTU              *int    `json:"mtu,omitempty" flag:"mtu"`
	LogOpts          *string `json:"logOpts,omitempty" flag:"log-opts"`
}

var validDockerConfigStorageValues = []string{"aufs", "btrfs", "devicemapper", "overlay", "overlay2", "zfs"}

func ValidateDockerConfig(config *DockerConfig, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, IsValidValue(fldPath.Child("storage"), config.Storage, validDockerConfigStorageValues)...)
	return allErrs
}
