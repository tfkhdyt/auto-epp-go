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
			typePath := filepath.Join(powerSupplyPath, file.Name(), "type")

			typeData, err := os.ReadFile(typePath)
			if err != nil {
				continue
			}

			supplyType := strings.TrimSpace(string(typeData))
			if supplyType == "Mains" {
				onlinePath := filepath.Join(powerSupplyPath, file.Name(), "online")

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
				statusPath := filepath.Join(powerSupplyPath, file.Name(), "status")
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
