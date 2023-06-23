package name

import (
	"HWIDSpoofer/internal/util"
	"fmt"
	"golang.org/x/sys/windows/registry"
)

const (
	logPrefix = "PC Name Spoofer"
	path1     = "SYSTEM\\CurrentControlSet\\Control\\ComputerName\\ComputerName"
	path2     = "SYSTEM\\CurrentControlSet\\Control\\ComputerName\\ActiveComputerName"
	path3     = "SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters"
	path4     = "SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters\\Interfaces"
)

func Spoof() {
	compName := util.RandSeq(8)

	k1, err := registry.OpenKey(registry.LOCAL_MACHINE, path1, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = k.Close()
	}(k1)

	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = k1.SetStringValue("ComputerName", compName)
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}
	err = k1.SetStringValue("ActiveComputerName", compName)
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}
	err = k1.SetStringValue("ComputerNamePhysicalDnsDomain", "")
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	k2, err := registry.OpenKey(registry.LOCAL_MACHINE, path2, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = k.Close()
	}(k2)

	err = k2.SetStringValue("ComputerName", compName)
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}
	err = k2.SetStringValue("ActiveComputerName", compName)
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}
	err = k2.SetStringValue("ComputerNamePhysicalDnsDomain", "")
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	k3, err := registry.OpenKey(registry.LOCAL_MACHINE, path3, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = k.Close()
	}(k3)

	err = k3.SetStringValue("Hostname", compName)
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}
	err = k3.SetStringValue("NV Hostname", compName)
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	k4s, err := registry.OpenKey(registry.LOCAL_MACHINE, path4, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = k.Close()
	}(k3)

	interfaceKeys, err := k4s.ReadSubKeyNames(0)
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	for _, interfaceKey := range interfaceKeys {
		path := fmt.Sprintf("%s\\%s", path4, interfaceKey)

		xx, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.ALL_ACCESS)

		err = xx.SetStringValue("Hostname", compName)
		if err != nil {
			util.Log(logPrefix, err.Error())
		}
		err = xx.SetStringValue("NV Hostname", compName)
		if err != nil {
			util.Log(logPrefix, err.Error())
		}

		err = xx.Close()
		if err != nil {
			util.Log(logPrefix, err.Error())
		}
	}

	util.Log(logPrefix, "Success")
}
