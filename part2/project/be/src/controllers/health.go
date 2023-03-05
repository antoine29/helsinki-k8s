package controllers

import (
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
	err := pg.IsDBHealthy()
	if err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "Healthy"})
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
}
