package service

import (
	"easygoadmin/dto"
	"easygoadmin/model"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"easygoadmin/vo"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var Dept = new(deptService)

type deptService struct{}

func (s *deptService) GetList(req dto.DeptPageReq) ([]model.Dept, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 部门名称
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询数据
	var list []model.Dept
	err := query.Find(&list)
	return list, err
}

func (s *deptService) Add(req dto.DeptAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Dept
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = gconv.Int(req.Type)
	entity.Pid = gconv.Int(req.Pid)
	entity.Sort = gconv.Int(req.Sort)
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	entity.Mark = 1
	// 插入记录
	return entity.Insert()
}

func (s *deptService) Update(req dto.DeptUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Dept{Id: gconv.Int(req.Id)}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	// 设置参数
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = gconv.Int(req.Type)
	entity.Pid = gconv.Int(req.Pid)
	entity.Sort = gconv.Int(req.Sort)
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	// 更新记录
	return entity.Update()
}

func (s *deptService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := model.Dept{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, err
		}
		return rows, nil
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			id, _ := strconv.Atoi(v)
			entity := &model.Dept{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

// 获取子级菜单
func (s *deptService) GetDeptTreeList() ([]*vo.DeptTreeNode, error) {
	var deptNode vo.DeptTreeNode
	// 查询列表
	list := make([]model.Dept, 0)
	utils.XormDb.Where("mark=1").Cols("id,name,pid").OrderBy("sort asc").Find(&list)
	makeDeptTree(list, &deptNode)
	return deptNode.Children, nil
}

//递归生成分类列表
func makeDeptTree(cate []model.Dept, tn *vo.DeptTreeNode) {
	for _, c := range cate {
		if c.Pid == tn.Id {
			child := &vo.DeptTreeNode{}
			child.Dept = c
			tn.Children = append(tn.Children, child)
			makeDeptTree(cate, child)
		}
	}
}

// 数据源转换
func (s *deptService) MakeList(data []*vo.DeptTreeNode) map[int]string {
	deptList := make(map[int]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			deptList[val.Id] = val.Name

			// 二级栏目
			for _, v := range val.Children {
				deptList[v.Id] = "|--" + v.Name

				// 三级栏目
				for _, vt := range v.Children {
					deptList[vt.Id] = "|--|--" + vt.Name
				}
			}
		}
	}
	return deptList
}
