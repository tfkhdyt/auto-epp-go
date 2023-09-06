package checker

import (
	"fmt"
	"os"
)

func CheckRoot() {
	if os.Geteuid() != 0 {
		fmt.Println("auto-epp-go must be run with root privileges.")
		os.Exit(1)
	}
}
