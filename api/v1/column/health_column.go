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

type HealthColumnApi struct {
}

var healthColumnService = service.ServiceGroupApp.ColumnServiceGroup.HealthColumnService

// CreateHealthColumn 创建HealthColumn
// @Tags HealthColumn
// @Summary 创建HealthColumn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body column.HealthColumn true "创建HealthColumn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthColumn/createHealthColumn [post]
func (healthColumnApi *HealthColumnApi) CreateHealthColumn(c *gin.Context) {
	var healthColumn column.HealthColumn
	err := c.ShouldBindJSON(&healthColumn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := healthColumnService.CreateHealthColumn(healthColumn); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHealthColumn 删除HealthColumn
// @Tags HealthColumn
// @Summary 删除HealthColumn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body column.HealthColumn true "删除HealthColumn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /healthColumn/deleteHealthColumn [delete]
func (healthColumnApi *HealthColumnApi) DeleteHealthColumn(c *gin.Context) {
	var healthColumn column.HealthColumn
	err := c.ShouldBindJSON(&healthColumn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := healthColumnService.DeleteHealthColumn(healthColumn); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHealthColumnByIds 批量删除HealthColumn
// @Tags HealthColumn
// @Summary 批量删除HealthColumn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HealthColumn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /healthColumn/deleteHealthColumnByIds [delete]
func (healthColumnApi *HealthColumnApi) DeleteHealthColumnByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := healthColumnService.DeleteHealthColumnByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHealthColumn 更新HealthColumn
// @Tags HealthColumn
// @Summary 更新HealthColumn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body column.HealthColumn true "更新HealthColumn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /healthColumn/updateHealthColumn [put]
func (healthColumnApi *HealthColumnApi) UpdateHealthColumn(c *gin.Context) {
	var healthColumn column.HealthColumn
	err := c.ShouldBindJSON(&healthColumn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := healthColumnService.UpdateHealthColumn(healthColumn); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHealthColumn 用id查询HealthColumn
// @Tags HealthColumn
// @Summary 用id查询HealthColumn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query column.HealthColumn true "用id查询HealthColumn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /healthColumn/findHealthColumn [get]
func (healthColumnApi *HealthColumnApi) FindHealthColumn(c *gin.Context) {
	var healthColumn column.HealthColumn
	err := c.ShouldBindQuery(&healthColumn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehealthColumn, err := healthColumnService.GetHealthColumn(healthColumn.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehealthColumn": rehealthColumn}, c)
	}
}

// GetHealthColumnList 分页获取HealthColumn列表
// @Tags HealthColumn
// @Summary 分页获取HealthColumn列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query columnReq.HealthColumnSearch true "分页获取HealthColumn列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /healthColumn/getHealthColumnList [get]
func (healthColumnApi *HealthColumnApi) GetHealthColumnList(c *gin.Context) {
	var pageInfo columnReq.HealthColumnSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := healthColumnService.GetHealthColumnInfoList(pageInfo); err != nil {
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
