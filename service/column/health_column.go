package column

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/column"
	columnReq "github.com/flipped-aurora/gin-vue-admin/server/model/column/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HealthColumnService struct {
}

// CreateHealthColumn 创建HealthColumn记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnService *HealthColumnService) CreateHealthColumn(healthColumn column.HealthColumn) (err error) {
	err = global.GVA_DB.Create(&healthColumn).Error
	return err
}

// DeleteHealthColumn 删除HealthColumn记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnService *HealthColumnService) DeleteHealthColumn(healthColumn column.HealthColumn) (err error) {
	err = global.GVA_DB.Delete(&healthColumn).Error
	return err
}

// DeleteHealthColumnByIds 批量删除HealthColumn记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnService *HealthColumnService) DeleteHealthColumnByIds(ids request.IdsReq, delete_by uint) (err error) {
	err = global.GVA_DB.Delete(&[]column.HealthColumn{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHealthColumn 更新HealthColumn记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnService *HealthColumnService) UpdateHealthColumn(healthColumn column.HealthColumn) (err error) {
	err = global.GVA_DB.Save(&healthColumn).Error
	return err
}

// GetHealthColumn 根据id获取HealthColumn记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnService *HealthColumnService) GetHealthColumn(id uint) (healthColumn column.HealthColumn, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&healthColumn).Error
	return
}

// GetHealthColumnInfoList 分页获取HealthColumn记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnService *HealthColumnService) GetHealthColumnInfoList(info columnReq.HealthColumnSearch) (list []column.HealthColumn, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&column.HealthColumn{})
	var healthColumns []column.HealthColumn
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&healthColumns).Error
	return healthColumns, total, err
}
