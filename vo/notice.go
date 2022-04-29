package vo

import "easygoadmin/model"

// 通知公告Vo
type NoticeInfoVo struct {
	model.Notice
	SourceName string `json:"sourceName"` // 通知来源
}
