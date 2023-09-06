package set

import (
	"fmt"
	"os"
	"runtime"
)

func SetGovernor() {
	cpuCount := runtime.NumCPU()

	for cpu := 0; cpu < cpuCount; cpu++ {
		governorFilePath := fmt.
			Sprintf("/sys/devices/system/cpu/cpu%d/cpufreq/scaling_governor", cpu)
		if err := os.WriteFile(
			governorFilePath,
			[]byte("powersave"),
			0644,
		); err != nil {
			fmt.Println("Failed to set scaling governor:", err)
			os.Exit(1)
		}
	}
}
