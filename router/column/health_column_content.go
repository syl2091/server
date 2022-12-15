package column

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HealthColumnContentRouter struct {
}

// InitHealthColumnContentRouter 初始化 HealthColumnContent 路由信息
func (s *HealthColumnContentRouter) InitHealthColumnContentRouter(Router *gin.RouterGroup) {
	healthColumnContentRouter := Router.Group("healthColumnContent").Use(middleware.OperationRecord())
	healthColumnContentRouterWithoutRecord := Router.Group("healthColumnContent")
	var healthColumnContentApi = v1.ApiGroupApp.ColumnApiGroup.HealthColumnContentApi
	{
		healthColumnContentRouter.POST("createHealthColumnContent", healthColumnContentApi.CreateHealthColumnContent)   // 新建HealthColumnContent
		healthColumnContentRouter.DELETE("deleteHealthColumnContent", healthColumnContentApi.DeleteHealthColumnContent) // 删除HealthColumnContent
		healthColumnContentRouter.DELETE("deleteHealthColumnContentByIds", healthColumnContentApi.DeleteHealthColumnContentByIds) // 批量删除HealthColumnContent
		healthColumnContentRouter.PUT("updateHealthColumnContent", healthColumnContentApi.UpdateHealthColumnContent)    // 更新HealthColumnContent
	}
	{
		healthColumnContentRouterWithoutRecord.GET("findHealthColumnContent", healthColumnContentApi.FindHealthColumnContent)        // 根据ID获取HealthColumnContent
		healthColumnContentRouterWithoutRecord.GET("getHealthColumnContentList", healthColumnContentApi.GetHealthColumnContentList)  // 获取HealthColumnContent列表
	}
}
