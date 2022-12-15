// 自动生成模板HealthColumn
package column

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HealthColumn 结构体
type HealthColumn struct {
      global.GVA_MODEL
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建日期;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新日期;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:栏目名称;size:32;"`
      Pid  *int `json:"pid" form:"pid" gorm:"column:pid;comment:父级id;size:10;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
      Type  string `json:"type" form:"type" gorm:"column:type;comment:类型;size:10;"`
      ImageLink  string `json:"imageLink" form:"imageLink" gorm:"column:image_link;comment:图片链接;size:100;"`
      State  string `json:"state" form:"state" gorm:"column:state;comment:状态;size:32;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
      HospitalOrgCode  string `json:"hospitalOrgCode" form:"hospitalOrgCode" gorm:"column:hospital_org_code;comment:医院组织代码;size:32;"`
      IsDel  string `json:"isDel" form:"isDel" gorm:"column:is_del;comment:删除状态;size:1;"`
}


// TableName HealthColumn 表名
func (HealthColumn) TableName() string {
  return "health_column"
}

