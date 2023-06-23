package main

import (
	biosSpoofer "HWIDSpoofer/internal/bios"
	efiSpoofer "HWIDSpoofer/internal/efi"
	guidSpoofer "HWIDSpoofer/internal/guid"
	installationSpoofer "HWIDSpoofer/internal/installation"
	macSpoofer "HWIDSpoofer/internal/mac"
	pcNameSpoofer "HWIDSpoofer/internal/name"
)

func main() {
	//k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion32232", registry.QUERY_VALUE)
	//
	//fmt.Println(k.GetStringValue("InstallationID"))
	//fmt.Println(err)

	installationSpoofer.Spoof()
	pcNameSpoofer.Spoof()
	guidSpoofer.Spoof()
	biosSpoofer.Spoof()
	efiSpoofer.Spoof()
	macSpoofer.Spoof()
}
