package mac

import (
	"HWIDSpoofer/internal/util"
	"fmt"
	"golang.org/x/sys/windows/registry"
)

const (
	logPrefix = "MAC Spoofer"
	path      = "SYSTEM\\CurrentControlSet\\Control\\Class"
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

	subs, err := k.ReadSubKeyNames(0)

	for _, l2 := range subs {
		l2path := fmt.Sprintf("%s\\%s", path, l2)

		l2k, err := registry.OpenKey(registry.LOCAL_MACHINE, l2path, registry.ALL_ACCESS)
		if err != nil {
			util.Log(logPrefix, err.Error())
			return
		}

		lows, err := l2k.ReadSubKeyNames(0)
		if err != nil {
			util.Log(logPrefix, err.Error())
			return
		}

		for _, l3 := range lows {
			if l3 == "Properties" {
				continue
			}

			l3path := fmt.Sprintf("%s\\%s", l2path, l3)

			l3k, err := registry.OpenKey(registry.LOCAL_MACHINE, l3path, registry.ALL_ACCESS)
			if err != nil {
				util.Log(logPrefix, err.Error())
				return
			}

			err = l3k.SetStringValue("NetworkAddress", util.GenMac())
			if err != nil {
				util.Log(logPrefix, err.Error())
				continue
			}

			err = l3k.Close()
			if err != nil {
				util.Log(logPrefix, err.Error())
				continue
			}
		}

		err = l2k.Close()
		if err != nil {
			util.Log(logPrefix, err.Error())
			return
		}
	}

	if err != nil {
		util.Log(logPrefix, err.Error())
		return
	}

	util.Log(logPrefix, "Success")
}
