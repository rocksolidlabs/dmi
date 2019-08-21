package dmi

import (
	"fmt"
)

type ManagementDeviceThresholdData struct {
	infoCommon
	LowerThresholdNonCritical    uint16
	UpperThresholdNonCritical    uint16
	LowerThresholdCritical       uint16
	UpperThresholdCritical       uint16
	LowerThresholdNonRecoverable uint16
	UpperThresholdNonRecoverable uint16
}

func (m ManagementDeviceThresholdData) String() string {
	return fmt.Sprintf("Management Device Threshold Data\n"+
		"\tLower Threshold Non Critical: %d\n"+
		"\tUpper Threshold Non Critical: %d\n"+
		"\tLower Threshold Critical: %d\n"+
		"\tUpper Threshold Critical: %d\n"+
		"\tLower Threshold Non Recoverable: %d\n"+
		"\tUpper Threshold Non Recoverable: %d",
		m.LowerThresholdNonCritical,
		m.UpperThresholdNonCritical,
		m.LowerThresholdCritical,
		m.UpperThresholdCritical,
		m.LowerThresholdNonRecoverable,
		m.UpperThresholdNonRecoverable)
}
