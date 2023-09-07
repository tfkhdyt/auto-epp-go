package config

const (
	configFile    = "/etc/auto-epp-go.conf"
	defaultConfig = `# see available epp state by running: cat /sys/devices/system/cpu/cpu0/cpufreq/energy_performance_available_preferences
[Settings]
epp_state_for_AC=balance_performance
epp_state_for_BAT=power
`
)
