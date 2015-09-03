package arch

import (
	"fmt"

	"../models"
	"./arm"
	"./arm64"
	"./mips"
	"./x86"
	"./x86_64"
)

var archMap = map[string]*models.Arch{
	"arm":    arm.Arch,
	"arm64":  arm64.Arch,
	"mips":   mips.Arch,
	"x86":    x86.Arch,
	"x86_64": x86_64.Arch,
}

func GetArch(name, os string) (*models.Arch, *models.OS, error) {
	a, ok := archMap[name]
	if !ok {
		return nil, nil, fmt.Errorf("Arch '%s' not found.", name)
	}
	o, ok := a.OS[os]
	if !ok {
		return nil, nil, fmt.Errorf("OS '%s' not found for arch '%s'.", os, name)
	}
	return a, o, nil
}
