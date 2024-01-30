package services

type IpConfig struct {
	IP       string
	Hostname string
	Active   bool
}

func getConfig() []IpConfig {
	// Mock implementation using sample data (replace with actual data source)
	return []IpConfig{
		{"127.0.0.1", "mta-prod-1", true},
		{"127.0.0.2", "mta-prod-1", false},
		{"127.0.0.3", "mta-prod-2", true},
		{"127.0.0.4", "mta-prod-2", true},
		{"127.0.0.5", "mta-prod-2", true},
		{"127.0.0.6", "mta-prod-3", false},
	}
}
