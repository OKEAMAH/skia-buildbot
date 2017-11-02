package main

import (
	"go.skia.org/infra/go/gce"
	"go.skia.org/infra/go/gce/server"
)

func Prod() *gce.Instance {
	vm := server.Server20170928("skia-leasing")
	vm.DataDisks[0].SizeGb = 100
	vm.DataDisks[0].MountPath = "/mnt/pd0"
	vm.DataDisks[0].Type = gce.DISK_TYPE_PERSISTENT_STANDARD
	vm.MachineType = gce.MACHINE_TYPE_HIGHMEM_2
	vm.Metadata["owner_primary"] = "rmistry"
	vm.Metadata["owner_secondary"] = "kjlubick"
	return vm
}

func main() {
	server.Main(gce.ZONE_DEFAULT, map[string]*gce.Instance{
		"prod": Prod(),
	})
}
