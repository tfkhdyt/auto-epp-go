package config

const (
	configFile    = "/etc/auto-epp.conf"
	defaultConfig = `# see available epp state by running: cat /sys/devices/system/cpu/cpu0/cpufreq/energy_performance_available_preferences
[Settings]
epp_state_for_AC=performance
epp_state_for_BAT=balance_power
`
)
