package column

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HealthColumnRouter struct {
}

// InitHealthColumnRouter 初始化 HealthColumn 路由信息
func (s *HealthColumnRouter) InitHealthColumnRouter(Router *gin.RouterGroup) {
	healthColumnRouter := Router.Group("healthColumn").Use(middleware.OperationRecord())
	healthColumnRouterWithoutRecord := Router.Group("healthColumn")
	var healthColumnApi = v1.ApiGroupApp.ColumnApiGroup.HealthColumnApi
	{
		healthColumnRouter.POST("createHealthColumn", healthColumnApi.CreateHealthColumn)   // 新建HealthColumn
		healthColumnRouter.DELETE("deleteHealthColumn", healthColumnApi.DeleteHealthColumn) // 删除HealthColumn
		healthColumnRouter.DELETE("deleteHealthColumnByIds", healthColumnApi.DeleteHealthColumnByIds) // 批量删除HealthColumn
		healthColumnRouter.PUT("updateHealthColumn", healthColumnApi.UpdateHealthColumn)    // 更新HealthColumn
	}
	{
		healthColumnRouterWithoutRecord.GET("findHealthColumn", healthColumnApi.FindHealthColumn)        // 根据ID获取HealthColumn
		healthColumnRouterWithoutRecord.GET("getHealthColumnList", healthColumnApi.GetHealthColumnList)  // 获取HealthColumn列表
	}
}
