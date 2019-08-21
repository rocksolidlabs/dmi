package dmi

import (
	"fmt"
)

type SystemConfigurationOptions struct {
	infoCommon
	Count   byte
	strings string
}

func (s SystemConfigurationOptions) String() string {
	return fmt.Sprintf("System Configuration Option\n\t\t%s", s.strings)
}

func newSystemConfigurationOptions(h dmiHeader) dmiTyper {
	var sc SystemConfigurationOptions
	data := h.data
	sc.Count = data[0x04]
	for i := byte(1); i <= sc.Count; i++ {
		sc.strings += fmt.Sprintf("string %d: %s\n\t\t", i, h.FieldString(int(data[0x04+i])))
	}
	return &sc
}

func GetSystemConfigurationOptions() *SystemConfigurationOptions {
	if d, ok := gdmi[SMBIOSStructureTypeSystemConfigurationOptions]; ok {
		return d.(*SystemConfigurationOptions)
	}
	return nil
}

func init() {
	addTypeFunc(SMBIOSStructureTypeSystemConfigurationOptions, newSystemConfigurationOptions)
}
