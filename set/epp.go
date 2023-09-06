package set

import (
	"fmt"
	"os"
	"runtime"
)

func SetEPP(eppValue string) {
	cpuCount := runtime.NumCPU()
	for cpu := 0; cpu < cpuCount; cpu++ {
		eppFilePath := fmt.Sprintf(
			"/sys/devices/system/cpu/cpu%d/cpufreq/energy_performance_preference",
			cpu,
		)

		if err := os.WriteFile(eppFilePath, []byte(eppValue), 0644); err != nil {
			fmt.Println("Failed to set energy performance preference:", err)
			os.Exit(1)
		}
	}
}
