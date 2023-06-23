package installation

import (
	"HWIDSpoofer/internal/util"
	"golang.org/x/sys/windows/registry"
)

const (
	logPrefix = "Installation Spoofer"
	tag       = "InstallationID"
	path      = "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion"
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

	err = k.SetStringValue(tag, util.GenGuid())
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	util.Log(logPrefix, "Success")
}
