package vo

import "easygoadmin/model"

// 栏目信息
type ItemCateInfoVo struct {
	model.ItemCate
	ItemName string `json:"itemName"` // 栏目名称
}

// 栏目树结构
type CateTreeNode struct {
	model.ItemCate
	Children []*CateTreeNode `json:"children"` // 子栏目
}
