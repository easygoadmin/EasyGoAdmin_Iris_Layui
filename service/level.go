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

var Level = new(levelService)

type levelService struct{}

func (s *levelService) GetList(req dto.LevelPageReq) ([]model.Level, int64, error) {
	// 初始化查询实例
	query := utils.XormDb.Where("mark=1")
	// 字典名称查询
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	// 排序
	query = query.Asc("sort")
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	list := make([]model.Level, 0)
	count, err := query.FindAndCount(&list)
	// 返回结果
	return list, count, err
}

func (s *levelService) Add(req dto.LevelAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Level
	entity.Name = req.Name
	entity.Status = gconv.Int(req.Status)
	entity.Sort = gconv.Int(req.Sort)
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *levelService) Update(req dto.LevelUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Level{Id: gconv.Int(req.Id)}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}
	entity.Name = req.Name
	entity.Status = gconv.Int(req.Status)
	entity.Sort = gconv.Int(req.Sort)
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	// 更新记录
	return entity.Update()
}

func (s *levelService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Level{Id: gconv.Int(ids)}
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
			entity := &model.Level{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *levelService) Status(req dto.LevelStatusReq, userId int) (int64, error) {
	// 查询记录是否存在
	info := &model.Level{Id: gconv.Int(req.Id)}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity := &model.Level{}
	entity.Id = info.Id
	entity.Status = gconv.Int(req.Status)
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	return entity.Update()
}
