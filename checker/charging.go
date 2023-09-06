package checker

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CheckChargingStatus() bool {
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
			typeData, err := os.ReadFile(filepath.Join(powerSupplyPath, file.Name(), "type"))
			if err != nil {
				continue
			}

			supplyType := strings.TrimSpace(string(typeData))
			if supplyType == "Mains" {
				onlineData, err := os.ReadFile(filepath.Join(powerSupplyPath, file.Name(), "online"))
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
				statusData, err := os.ReadFile(filepath.Join(powerSupplyPath, file.Name(), "status"))
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
