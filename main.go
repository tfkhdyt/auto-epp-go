package main

import (
	"time"

	"codeberg.org/tfkhdyt/auto-epp-go/checker"
	"codeberg.org/tfkhdyt/auto-epp-go/config"
	"codeberg.org/tfkhdyt/auto-epp-go/set"
)

func main() {
	checker.CheckRoot()
	checker.CheckDriver()

	eppStateForAC, eppStateForBAT := config.ReadConfig()

	for {
		set.SetGovernor()

		if checker.CheckChargingStatus() {
			set.SetEPP(eppStateForAC)
		} else {
			set.SetEPP(eppStateForBAT)
		}

		time.Sleep(2 * time.Second)
	}
}
