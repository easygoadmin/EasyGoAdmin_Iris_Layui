package vo

import "easygoadmin/model"

// 广告信息Vo
type AdInfoVo struct {
	model.Ad
	TypeName   string `json:"typeName"`   // 广告类型
	AdSortDesc string `json:"adSortDesc"` // 广告位描述
}
