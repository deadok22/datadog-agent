// Code generated by go generate; DO NOT EDIT.
// +build linux_bpf

package runtime

import (
	"github.com/DataDog/datadog-agent/pkg/ebpf"
)

var Tracer = ebpf.NewRuntimeAsset("tracer.c", "ea7b70b93bf5605ec4a61bd72e64623de3a648ba8557aff2446d9e19076d150d")
