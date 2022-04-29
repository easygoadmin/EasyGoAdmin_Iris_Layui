package vo

import "easygoadmin/model"

// 配置数据列表
type ConfigDataVo struct {
	model.ConfigData
	TypeName string `json:"typeName"`
}
