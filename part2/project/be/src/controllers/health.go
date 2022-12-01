package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	pg "antoine29/go/web-server/src/dao/pg"
)

// @Summary Check API health
// @Schemes http
// @Description Check API health
// @Produce json
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	healthy, errorMessagePointer := pg.IsDBHealthy()
	if healthy {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "Healthy"})
		return
	}

	fmt.Println(errorMessagePointer)
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": errorMessagePointer})
}
