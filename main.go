package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	CONFIG_FILE    = "/etc/auto-epp.conf"
	DEFAULT_CONFIG = `# see available epp state by running: cat /sys/devices/system/cpu/cpu0/cpufreq/energy_performance_available_preferences
[Settings]
epp_state_for_AC=balance_performance
epp_state_for_BAT=power
`
)

func checkRoot() {
	if os.Geteuid() != 0 {
		fmt.Println("auto-epp must be run with root privileges.")
		os.Exit(1)
	}
}

func checkDriver() {
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

func readConfig() (string, string) {
	if _, err := os.Stat(CONFIG_FILE); os.IsNotExist(err) {
		err := os.WriteFile(CONFIG_FILE, []byte(DEFAULT_CONFIG), 0644)
		if err != nil {
			fmt.Println("Failed to create config file:", err)
			os.Exit(1)
		}
	}
	configData, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		os.Exit(1)
	}
	configStr := string(configData)
	configLines := strings.Split(configStr, "\n")
	eppStateForAC := ""
	eppStateForBAT := ""
	for _, line := range configLines {
		if strings.HasPrefix(line, "epp_state_for_AC") {
			eppStateForAC = strings.Split(line, "=")[1]
		} else if strings.HasPrefix(line, "epp_state_for_BAT") {
			eppStateForBAT = strings.Split(line, "=")[1]
		}
	}
	return eppStateForAC, eppStateForBAT
}

func charging() bool {
	powerSupplyPath := "/sys/class/power_supply/"
	files, err := os.ReadDir(powerSupplyPath)
	if err != nil {
		fmt.Println("Failed to read power supply directory:", err)
		os.Exit(1)
	}
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			continue
		}
		if file.IsDir() || (fileInfo.Mode()&os.ModeSymlink != 0) {
			typePath := powerSupplyPath + file.Name() + "/type"
			typeData, err := os.ReadFile(typePath)
			if err != nil {
				continue
			}
			supplyType := strings.TrimSpace(string(typeData))
			if supplyType == "Mains" {
				onlinePath := powerSupplyPath + file.Name() + "/online"
				onlineData, err := os.ReadFile(onlinePath)
				if err != nil {
					continue
				}
				val, err := strconv.Atoi(strings.TrimSpace(string(onlineData)))
				if err != nil {
					continue
				}
				if val == 1 {
					return true
				}
			} else if supplyType == "Battery" {
				statusPath := powerSupplyPath + file.Name() + "/status"
				statusData, err := os.ReadFile(statusPath)
				if err != nil {
					continue
				}
				status := strings.TrimSpace(string(statusData))
				if status == "Discharging" {
					return false
				}
			}
		}
	}
	return true
}

func setGovernor() {
	cpuCount := runtime.NumCPU()
	for cpu := 0; cpu < cpuCount; cpu++ {
		governorFilePath := fmt.Sprintf("/sys/devices/system/cpu/cpu%d/cpufreq/scaling_governor", cpu)
		err := os.WriteFile(governorFilePath, []byte("powersave"), 0644)
		if err != nil {
			fmt.Println("Failed to set scaling governor:", err)
			os.Exit(1)
		}
	}
}

func setEPP(eppValue string) {
	cpuCount := runtime.NumCPU()
	for cpu := 0; cpu < cpuCount; cpu++ {
		eppFilePath := fmt.Sprintf("/sys/devices/system/cpu/cpu%d/cpufreq/energy_performance_preference", cpu)
		err := os.WriteFile(eppFilePath, []byte(eppValue), 0644)
		if err != nil {
			fmt.Println("Failed to set energy performance preference:", err)
			os.Exit(1)
		}
	}
}

func main() {
	checkRoot()
	checkDriver()
	eppStateForAC, eppStateForBAT := readConfig()
	for {
		setGovernor()
		if charging() {
			setEPP(eppStateForAC)
		} else {
			setEPP(eppStateForBAT)
		}
		time.Sleep(2 * time.Second)
	}
}
