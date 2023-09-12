package apiv1

import (
	"github.com/netsepio/netsepio-gateway/api/v1/delegatereviewsubmission"
	"github.com/netsepio/netsepio-gateway/api/v1/status"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		// profile.ApplyRoutes(v1)
		delegatereviewsubmission.ApplyRoutes(v1)
		status.ApplyRoutes(v1)
	}
}
