package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mta-hosting-optimizer/services"
	"net/http"
	"os"
)

func GetHostnames(c *gin.Context) {
	thresholdStr := os.Getenv("THRESHOLD")
	threshold := 1 // Default threshold

	if thresholdStr != "" {
		fmt.Sscanf(thresholdStr, "%d", &threshold)
	}

	worker := services.PopulateGetInactiveHostsWorker()
	worker.DataService.PopulateData()
	hostnames := worker.GetInactiveHostnames(threshold)
	c.JSON(http.StatusOK, hostnames)
}
