package services

var ipConfigData = getConfig()

func GetInactiveHostnames(threshold int) []string {
	hostnames := make(map[string]int)

	for _, config := range ipConfigData {
		if !config.Active {
			hostnames[config.Hostname]++
		}
	}

	result := make([]string, 0)
	for hostname, count := range hostnames {
		if count <= threshold {
			result = append(result, hostname)
		}
	}

	return result
}
