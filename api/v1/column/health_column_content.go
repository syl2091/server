package column

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/column"
	columnReq "github.com/flipped-aurora/gin-vue-admin/server/model/column/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HealthColumnContentApi struct {
}

var healthColumnContentService = service.ServiceGroupApp.ColumnServiceGroup.HealthColumnContentService

// CreateHealthColumnContent 创建HealthColumnContent
// @Tags HealthColumnContent
// @Summary 创建HealthColumnContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body column.HealthColumnContent true "创建HealthColumnContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthColumnContent/createHealthColumnContent [post]
func (healthColumnContentApi *HealthColumnContentApi) CreateHealthColumnContent(c *gin.Context) {
	var healthColumnContent column.HealthColumnContent
	err := c.ShouldBindJSON(&healthColumnContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := healthColumnContentService.CreateHealthColumnContent(healthColumnContent); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHealthColumnContent 删除HealthColumnContent
// @Tags HealthColumnContent
// @Summary 删除HealthColumnContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body column.HealthColumnContent true "删除HealthColumnContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthColumnContent/deleteHealthColumnContent [delete]
func (healthColumnContentApi *HealthColumnContentApi) DeleteHealthColumnContent(c *gin.Context) {
	var healthColumnContent column.HealthColumnContent
	err := c.ShouldBindJSON(&healthColumnContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := healthColumnContentService.DeleteHealthColumnContent(healthColumnContent); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHealthColumnContentByIds 批量删除HealthColumnContent
// @Tags HealthColumnContent
// @Summary 批量删除HealthColumnContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HealthColumnContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /healthColumnContent/deleteHealthColumnContentByIds [delete]
func (healthColumnContentApi *HealthColumnContentApi) DeleteHealthColumnContentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := healthColumnContentService.DeleteHealthColumnContentByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHealthColumnContent 更新HealthColumnContent
// @Tags HealthColumnContent
// @Summary 更新HealthColumnContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body column.HealthColumnContent true "更新HealthColumnContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /healthColumnContent/updateHealthColumnContent [put]
func (healthColumnContentApi *HealthColumnContentApi) UpdateHealthColumnContent(c *gin.Context) {
	var healthColumnContent column.HealthColumnContent
	err := c.ShouldBindJSON(&healthColumnContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := healthColumnContentService.UpdateHealthColumnContent(healthColumnContent); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHealthColumnContent 用id查询HealthColumnContent
// @Tags HealthColumnContent
// @Summary 用id查询HealthColumnContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query column.HealthColumnContent true "用id查询HealthColumnContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /healthColumnContent/findHealthColumnContent [get]
func (healthColumnContentApi *HealthColumnContentApi) FindHealthColumnContent(c *gin.Context) {
	var healthColumnContent column.HealthColumnContent
	err := c.ShouldBindQuery(&healthColumnContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehealthColumnContent, err := healthColumnContentService.GetHealthColumnContent(healthColumnContent.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehealthColumnContent": rehealthColumnContent}, c)
	}
}

// GetHealthColumnContentList 分页获取HealthColumnContent列表
// @Tags HealthColumnContent
// @Summary 分页获取HealthColumnContent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query columnReq.HealthColumnContentSearch true "分页获取HealthColumnContent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthColumnContent/getHealthColumnContentList [get]
func (healthColumnContentApi *HealthColumnContentApi) GetHealthColumnContentList(c *gin.Context) {
	var pageInfo columnReq.HealthColumnContentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := healthColumnContentService.GetHealthColumnContentInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
