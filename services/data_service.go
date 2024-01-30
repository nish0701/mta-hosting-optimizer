package services

//go:generate mockgen  -destination=./mocked_services.go  -package=services . DataService

type DataService interface {
	PopulateData()
	GetData() []IpConfig
}

type DataWorker struct {
	populatedData []IpConfig
}

func PopulateDataWorker() *DataWorker {
	return &DataWorker{}
}

func (w *DataWorker) PopulateData() {
	w.populatedData = []IpConfig{
		{"127.0.0.1", "mta-prod-1", true},
		{"127.0.0.2", "mta-prod-1", false},
		{"127.0.0.3", "mta-prod-2", true},
		{"127.0.0.4", "mta-prod-2", true},
		{"127.0.0.5", "mta-prod-2", true},
		{"127.0.0.6", "mta-prod-3", false},
		{"127.0.0.6", "mta-prod-3", false},
	}
}

func (w *DataWorker) GetData() []IpConfig {
	return w.populatedData
}
