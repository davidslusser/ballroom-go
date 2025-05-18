package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/calculate-partners", CalculatePartners)

	router.Run(":8000") // Listen and serve on localhost:8000
}
