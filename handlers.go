package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CalculatePartners handles the /calculate-partners POST request
func CalculatePartners(c *gin.Context) {
	var req CalculationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if req.TotalLeaders != len(req.LeaderKnowledge) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "total_leaders does not match leader_knowledge size"})
		return
	}
	if req.TotalFollowers != len(req.FollowerKnowledge) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "total_followers does not match follower_knowledge size"})
		return
	}

	average := SimulateEvent(req)
	c.JSON(http.StatusOK, CalculationResponse{AverageDancePartners: average})
}
