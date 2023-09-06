package checker

import (
	"fmt"
	"os"
	"strings"
)

func CheckDriver() {
	scalingDriverPath := "/sys/devices/system/cpu/cpu0/cpufreq/scaling_driver"

	scalingDriver, err := os.ReadFile(scalingDriverPath)
	if err != nil {
		fmt.Println("The system is not running amd-pstate-epp")
		os.Exit(1)
	}

	if strings.TrimSpace(string(scalingDriver)) != "amd-pstate-epp" {
		fmt.Println("The system is not running amd-pstate-epp")
		os.Exit(1)
	}
}
