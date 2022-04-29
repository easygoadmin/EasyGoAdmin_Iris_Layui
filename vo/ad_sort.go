package vo

import "easygoadmin/model"

// 广告位信息
type AdSortInfoVo struct {
	model.AdSort
	ItemName     string `json:"itemName"`     // 站点名称
	CateName     string `json:"cateName"`     // 栏目名称
	PlatformName string `json:"platformName"` // 所属平台
}
