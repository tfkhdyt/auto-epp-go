package main

import (
	"time"

	"github.com/tfkhdyt/auto-epp-go/checker"
	"github.com/tfkhdyt/auto-epp-go/config"
	"github.com/tfkhdyt/auto-epp-go/set"
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
