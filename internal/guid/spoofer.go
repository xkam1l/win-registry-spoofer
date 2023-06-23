package guid

import (
	"HWIDSpoofer/internal/util"
	"fmt"
	"golang.org/x/sys/windows/registry"
)

const (
	logPrefix = "GUID Spoofer"
	path1     = "SYSTEM\\CurrentControlSet\\Control\\IDConfigDB\\Hardware Profiles\\0001"
	path2     = "SOFTWARE\\Microsoft\\Cryptography"
	path3     = "SOFTWARE\\Microsoft\\SQMClient"
	path4     = "SYSTEM\\CurrentControlSet\\Control\\SystemInformation"
	path5     = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\WindowsUpdate"
)

func Spoof() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path1, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = k.Close()
	}(k)

	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = k.SetStringValue("HwProfileGuid", fmt.Sprintf("%s", util.GenGuid()))
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	k2, err := registry.OpenKey(registry.LOCAL_MACHINE, path2, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = k2.Close()
	}(k2)

	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = k2.SetStringValue("MachineGuid", util.GenGuid())
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	k3, err := registry.OpenKey(registry.LOCAL_MACHINE, path3, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = k3.Close()
	}(k3)

	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = k3.SetStringValue("MachineId", fmt.Sprintf("%s", util.GenGuid()))
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	sysk, err := registry.OpenKey(registry.LOCAL_MACHINE, path4, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = sysk.Close()
	}(sysk)

	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	date := util.RandDate()

	err = sysk.SetStringValue("BIOSReleaseDate", fmt.Sprintf("%d/%d/%d", date.Day(), date.Month(), date.Year()))
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = sysk.SetStringValue("BIOSVersion", util.RandSeq(10))
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = sysk.SetStringValue("ComputerHardwareId", util.GenGuid())
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = sysk.SetStringValue("ComputerHardwareIds", fmt.Sprintf("%s\n%s\n%s", util.GenGuid(), util.GenGuid(), util.GenGuid()))
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = sysk.SetStringValue("SystemManufacturer", util.RandSeq(15))
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = sysk.SetStringValue("SystemProductName", util.RandSeq(6))
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	upk, err := registry.OpenKey(registry.LOCAL_MACHINE, path5, registry.ALL_ACCESS)
	defer func(k registry.Key) {
		_ = upk.Close()
	}(upk)

	err = upk.SetStringValue("SusClientId", util.GenGuid())
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	err = upk.SetBinaryValue("SusClientIdValidation", []byte(util.RandSeq(25)))
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	util.Log(logPrefix, "Success")
}
