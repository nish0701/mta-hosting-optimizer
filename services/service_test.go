package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ipConfigTestData struct {
	inputConfig    []IpConfig
	inputThreshold int
	expectedLength int
}

func populateTestData() []ipConfigTestData {
	return []ipConfigTestData{
		{
			inputConfig: []IpConfig{
				{"127.0.0.1", "mta-prod-1", true},
				{"127.0.0.2", "mta-prod-1", false},
				{"127.0.0.3", "mta-prod-2", true},
				{"127.0.0.4", "mta-prod-2", true},
				{"127.0.0.5", "mta-prod-2", true},
				{"127.0.0.6", "mta-prod-3", false},
			},
			inputThreshold: 1,
			expectedLength: 2,
		},
		{
			inputConfig: []IpConfig{
				{"127.0.0.1", "mta-prod-1", false},
				{"127.0.0.2", "mta-prod-1", false},
				{"127.0.0.3", "mta-prod-2", true},
				{"127.0.0.4", "mta-prod-2", true},
				{"127.0.0.5", "mta-prod-2", true},
				{"127.0.0.6", "mta-prod-3", false},
			},
			inputThreshold: 2,
			expectedLength: 1,
		},
		{
			inputConfig: []IpConfig{
				{"127.0.0.1", "mta-prod-1", false},
				{"127.0.0.2", "mta-prod-1", false},
				{"127.0.0.3", "mta-prod-2", true},
				{"127.0.0.4", "mta-prod-2", true},
				{"127.0.0.5", "mta-prod-2", true},
				{"127.0.0.6", "mta-prod-3", false},
				{"127.0.0.6", "mta-prod-3", false},
			},
			inputThreshold: 2,
			expectedLength: 2,
		},
		{
			inputConfig: []IpConfig{
				{"127.0.0.1", "mta-prod-1", false},
				{"127.0.0.2", "mta-prod-1", false},
				{"127.0.0.2", "mta-prod-1", false},
				{"127.0.0.3", "mta-prod-2", true},
				{"127.0.0.4", "mta-prod-2", true},
				{"127.0.0.5", "mta-prod-2", true},
				{"127.0.0.6", "mta-prod-3", false},
				{"127.0.0.6", "mta-prod-3", false},
				{"127.0.0.6", "mta-prod-3", false},
				{"127.0.0.6", "mta-prod-3", false},
			},
			inputThreshold: 3,
			expectedLength: 2,
		},
	}
}

func TestGetInactiveHostsWorker_GetInactiveHostnames(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedDataService := NewMockDataService(ctrl)
	worker := GetInactiveHostsWorker{
		mockedDataService,
	}
	testData := populateTestData()
	for i := 0; i < len(testData); i++ {
		name := fmt.Sprintf("test data %d", i)
		t.Run(name, func(t *testing.T) {
			mockedDataService.EXPECT().GetData().Return(testData[i].inputConfig).Times(1)
			mockedDataService.EXPECT().PopulateData().Times(0)
			result := worker.GetInactiveHostnames(testData[i].inputThreshold)
			assert.EqualValues(t, testData[i].expectedLength, len(result))
		})
	}
}

func TestDataWorker_PopulateData(t *testing.T) {
	w := PopulateDataWorker()
	w.PopulateData()
	data := w.GetData()
	assert.EqualValues(t, len(data), len(w.populatedData))
}
