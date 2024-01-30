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
		fmt.Println(thresholdStr, "%d", &threshold)
	}

	hostnames := services.GetInactiveHostnames(threshold)
	c.JSON(http.StatusOK, hostnames)
}
