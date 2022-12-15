package column

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/column"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    columnReq "github.com/flipped-aurora/gin-vue-admin/server/model/column/request"
)

type HealthColumnContentService struct {
}

// CreateHealthColumnContent 创建HealthColumnContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnContentService *HealthColumnContentService) CreateHealthColumnContent(healthColumnContent column.HealthColumnContent) (err error) {
	err = global.GVA_DB.Create(&healthColumnContent).Error
	return err
}

// DeleteHealthColumnContent 删除HealthColumnContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnContentService *HealthColumnContentService)DeleteHealthColumnContent(healthColumnContent column.HealthColumnContent) (err error) {
	err = global.GVA_DB.Delete(&healthColumnContent).Error
	return err
}

// DeleteHealthColumnContentByIds 批量删除HealthColumnContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnContentService *HealthColumnContentService)DeleteHealthColumnContentByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Delete(&[]column.HealthColumnContent{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHealthColumnContent 更新HealthColumnContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnContentService *HealthColumnContentService)UpdateHealthColumnContent(healthColumnContent column.HealthColumnContent) (err error) {
	err = global.GVA_DB.Save(&healthColumnContent).Error
	return err
}

// GetHealthColumnContent 根据id获取HealthColumnContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnContentService *HealthColumnContentService)GetHealthColumnContent(id uint) (healthColumnContent column.HealthColumnContent, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&healthColumnContent).Error
	return
}

// GetHealthColumnContentInfoList 分页获取HealthColumnContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthColumnContentService *HealthColumnContentService)GetHealthColumnContentInfoList(info columnReq.HealthColumnContentSearch) (list []column.HealthColumnContent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&column.HealthColumnContent{})
    var healthColumnContents []column.HealthColumnContent
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&healthColumnContents).Error
	return  healthColumnContents, total, err
}
