package status

import (
	"github.com/gin-gonic/gin"
	"github.com/netsepio/netsepio-gateway/api/types"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/status")
	{
		g.GET("", status)
	}
}

func status(c *gin.Context) {
	status := types.ApiResponse{
		Status:  200,
		Message: "LIVE",
	}
	c.JSON(200, status)
}
