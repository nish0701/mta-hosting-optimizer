package services

//var ipConfigData2 = getConfig()

type GetInactiveHostsService interface {
	GetInactiveHostnames(threshold int) []string
}

type GetInactiveHostsWorker struct {
	DataService
}

func PopulateGetInactiveHostsWorker() *GetInactiveHostsWorker {
	return &GetInactiveHostsWorker{
		PopulateDataWorker(),
	}
}

func (w *GetInactiveHostsWorker) GetInactiveHostnames(threshold int) []string {
	hostnames := make(map[string]int)
	ipConfigData := w.DataService.GetData()
	for _, config := range ipConfigData {
		if !config.Active {
			hostnames[config.Hostname]++
		}
	}

	result := make([]string, 0)
	for hostname, count := range hostnames {
		if count >= threshold {
			result = append(result, hostname)
		}
	}

	return result
}
