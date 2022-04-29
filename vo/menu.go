package vo

import "easygoadmin/model"

// 菜单Vo
type MenuTreeNode struct {
	model.Menu
	Children []*MenuTreeNode `json:"children"` // 子菜单
}
