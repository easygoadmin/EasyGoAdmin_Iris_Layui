package service

import (
	"easygoadmin/dto"
	"easygoadmin/model"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"strconv"
	"strings"
	"time"
)

var Dict = new(dictService)

type dictService struct{}

func (s *dictService) GetList(req dto.DictPageReq) ([]model.Dict, int64, error) {
	// 初始化查询实例
	query := utils.XormDb.Where("mark=1")
	// 职级名称查询
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	// 排序
	query = query.Asc("sort")
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	list := make([]model.Dict, 0)
	count, err := query.FindAndCount(&list)
	// 返回结果
	return list, count, err
}

func (s *dictService) Add(req dto.DictAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Dict
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	entity.Mark = 1

	// 插入记录
	return entity.Insert()
}

func (s *dictService) Update(req dto.DictUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Dict{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	if entity == nil {
		return 0, errors.New("记录不存在")
	}

	// 设置对象
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	// 更新记录
	return entity.Update()
}

func (s *dictService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Dict{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			id, _ := strconv.Atoi(v)
			entity := &model.Dict{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}
