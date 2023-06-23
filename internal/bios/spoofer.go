package bios

import (
	"HWIDSpoofer/internal/util"
	"golang.org/x/sys/windows/registry"
)

const (
	logPrefix = "GUID Spoofer"
	path      = "HARDWARE\\DESCRIPTION\\System\\BIOS"
)

func Spoof() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = k.Close()
	}(k)

	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = k.SetStringValue("SystemSerialNumber", util.RandSeq(10))
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	util.Log(logPrefix, "Success")
}
