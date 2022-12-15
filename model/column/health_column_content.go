// 自动生成模板HealthColumnContent
package column

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HealthColumnContent 结构体
type HealthColumnContent struct {
	global.GVA_MODEL
	CreateTime      *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建日期;"`
	UpdateTime      *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新日期;"`
	ColumnId        *int       `json:"columnId" form:"columnId" gorm:"column:column_id;comment:栏目id;size:10;"`
	ContentTitle    string     `json:"contentTitle" form:"contentTitle" gorm:"column:content_title;comment:内容标题;size:50;"`
	DisplayForm     string     `json:"displayForm" form:"displayForm" gorm:"column:display_form;comment:展现形式;size:32;"`
	Content         string     `json:"content" form:"content" gorm:"column:content;comment:内容;size:4294967295;"`
	ContentOne      string     `json:"contentOne" form:"contentOne" gorm:"column:content_one;comment:内容1;size:4294967295;"`
	ContentTwo      string     `json:"contentTwo" form:"contentTwo" gorm:"column:content_two;comment:内容2;size:4294967295;"`
	ImageOne        string     `json:"imageOne" form:"imageOne" gorm:"column:image_one;comment:图片1;size:1000;"`
	ImageTwo        string     `json:"imageTwo" form:"imageTwo" gorm:"column:image_two;comment:图片2;size:1000;"`
	VideoAssetsId   string     `json:"videoAssetsId" form:"videoAssetsId" gorm:"column:video_assets_id;comment:视频资产id;size:1000;"`
	State           string     `json:"state" form:"state" gorm:"column:state;comment:状态;size:1;"`
	IsDel           string     `json:"isDel" form:"isDel" gorm:"column:is_del;comment:删除状态;size:1;"`
	HospitalOrgCode string     `json:"hospitalOrgCode" form:"hospitalOrgCode" gorm:"column:hospital_org_code;comment:医院组织代码;size:32;"`
	Url             string     `json:"url" form:"url" gorm:"column:url;comment:外部链接url;size:1000;"`
	Sort            *int       `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
	IsTop           *bool      `json:"isTop" form:"isTop" gorm:"column:is_top;comment:是否置顶;"`
	ReleaseTime     *time.Time `json:"releaseTime" form:"releaseTime" gorm:"column:release_time;comment:发布时间;"`
	DeletedBy       uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName HealthColumnContent 表名
func (HealthColumnContent) TableName() string {
	return "health_column_content"
}
