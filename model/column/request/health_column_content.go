package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/column"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HealthColumnContentSearch struct{
    column.HealthColumnContent
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
