package vo

import "easygoadmin/model"

// 会员信息Vo
type MemberInfoVo struct {
	model.Member
	GenderName string      `json:"genderName"` // 性别
	DeviceName string      `json:"deviceName"` // 设备类型
	SourceName string      `json:"sourceName"` // 会员来源
	City       interface{} `json:"city"`       // 省市区
	CityName   string      `json:"cityName"`   // 城市名称
}
