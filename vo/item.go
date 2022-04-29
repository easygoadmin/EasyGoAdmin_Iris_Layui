package vo

import "easygoadmin/model"

// 站点信息Vo
type ItemInfoVo struct {
	model.Item
	TypeName string `json:"typeName"` // 站点类型
}
