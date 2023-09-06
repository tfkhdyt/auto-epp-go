package config

import (
	"fmt"
	"os"
	"strings"
)

func ReadConfig() (string, string) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err := os.WriteFile(configFile, []byte(defaultConfig), 0644)
		if err != nil {
			fmt.Println("Failed to create config file:", err)
			os.Exit(1)
		}
	}

	configData, err := os.ReadFile(configFile)
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
