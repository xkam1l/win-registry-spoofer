package efi

import (
	"HWIDSpoofer/internal/util"
	"fmt"
	"golang.org/x/sys/windows/registry"
)

const (
	logPrefix = "EFI Spoofer"
	path      = "SYSTEM\\CurrentControlSet\\Control\\Nsi"
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

	topKeys, err := k.ReadSubKeyNames(0)
	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	for _, topKey := range topKeys {
		mKeyPath := fmt.Sprintf("%s\\%s", path, topKey)

		mk, err := registry.OpenKey(registry.LOCAL_MACHINE, mKeyPath, registry.ALL_ACCESS)
		if err != nil {
			util.Log(logPrefix, err.Error())
			return
		}

		lowKeys, err := mk.ReadSubKeyNames(0)
		if err != nil {
			util.Log(logPrefix, err.Error())
			return
		}

		for _, lowKey := range lowKeys {
			lKeyPath := fmt.Sprintf("%s\\%s", mKeyPath, lowKey)
			lk, err := registry.OpenKey(registry.LOCAL_MACHINE, lKeyPath, registry.ALL_ACCESS)

			if err != nil {
				//util.Log(logPrefix, err.Error())
				//util.Log(logPrefix, "Skipping")

				// some will throw access error (about 5%, reason?)
				continue
			}

			err = lk.SetStringValue("VariableId", util.GenGuid())
			if err != nil {
				util.Log(logPrefix, err.Error())
				return
			}

			err = lk.Close()
			if err != nil {
				util.Log(logPrefix, err.Error())
				return
			}
		}

		err = mk.Close()
		if err != nil {
			util.Log(logPrefix, err.Error())
			return
		}
	}

	util.Log(logPrefix, "Success")
}
